package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/pythinh/go-news/internal/app/controller"
	"github.com/pythinh/go-news/internal/pkg/env"
	"github.com/pythinh/go-news/internal/pkg/types"
)

func main() {
	var conf types.Server
	env.Load(&conf)

	conns := controller.InitDB(&conf)
	defer conns.Close()

	log.Println("initializing HTTP routing...")
	routes, err := controller.InitRoute(conns)
	if err != nil {
		log.Panicln("failed to init routing, err:", err)
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
			log.Panicln("http.ListenAndServe() error:", err)
		}
	}()
	log.Println("HTTP server is listening at:", addr)

	// gracefully shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), conf.HTTP.ShutdownTimeout)
	defer cancel()
	log.Println("shutting down HTTP server...")
	err = httpServer.Shutdown(ctx)
	if err != nil {
		log.Panicln("HTTP server shutdown with error:", err)
	}
}
