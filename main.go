package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pythinh/go-news/internal/app/router"
	"github.com/pythinh/go-news/internal/pkg/env"
)

type srvConfig struct {
	HTTP struct {
		Address           string        `env:"HTTP_ADDRESS" default:""`
		Port              int           `env:"PORT" default:"8080"`
		ReadTimeout       time.Duration `env:"HTTP_READ_TIMEOUT" default:"5m"`
		WriteTimeout      time.Duration `env:"HTTP_WRITE_TIMEOUT" default:"5m"`
		ReadHeaderTimeout time.Duration `env:"HTTP_READ_HEADER_TIMEOUT" default:"30s"`
		ShutdownTimeout   time.Duration `env:"HTTP_SHUTDOWN_TIMEOUT" default:"10s"`
	}
}

func main() {
	var conf srvConfig
	env.Load(&conf)

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
