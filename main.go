package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

const gracefulShutdownPeriod = 10 * time.Second

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	cfg, err := NewConfig(ctx)
	if err != nil {
		log.Panicln("failed to load config:", err)
	}

	mcli, err := NewClient(ctx, cfg)
	if err != nil {
		log.Panicln("failed to create mongo client:", err)
	}
	defer func() {
		if err := mcli.Disconnect(ctx); err != nil {
			log.Panicln("failed to disconnect from mongo:", err)
		}
	}()

	handler := NewHandler(mcli)
	srv := NewServer(cfg, handler)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicln("failed to listen and serve:", err)
		}
	}()

	<-ctx.Done()
	tctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownPeriod)
	defer cancel()
	_ = tctx
	if err := srv.Shutdown(tctx); err != nil {
		log.Panicln("failed to shutdown server", err)
	}
	log.Println("server shutdown gracefully")
}
