package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/pythinh/go-news/internal/app/router"
	"github.com/pythinh/go-news/internal/app/types"
	"github.com/pythinh/go-news/internal/pkg/db"
	"github.com/pythinh/go-news/internal/pkg/env"
)

func main() {
	var conf types.Server
	env.Load(&conf)

	conns := db.Init(&conf)
	defer conns.Close()

	log.Println("initializing HTTP routing...")
	routes, err := router.Init()
	if err != nil {
		log.Panic("failed to init routing, err: ", err)
	}

	addr := fmt.Sprintf("%s:%d", conf.HTTP.Address, conf.HTTP.Port)
	httpServer := http.Server{
		Addr:              addr,
		Handler:           routes,
		ReadTimeout:       conf.HTTP.ReadTimeout,
		WriteTimeout:      conf.HTTP.WriteTimeout,
		ReadHeaderTimeout: conf.HTTP.ReadHeaderTimeout,
	}

	log.Println("starting HTTP server...")
	go func() {
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Panic("http.ListenAndServe() error: ", err)
		}
	}()
	log.Println("HTTP server is listening at: ", addr)

	// gracefully shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), conf.HTTP.ShutdownTimeout)
	defer cancel()
	log.Println("shutting down HTTP server...")
	err = httpServer.Shutdown(ctx)
	if err != nil {
		log.Panic("HTTP server shutdown with error: ", err)
	}
}
