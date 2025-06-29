package main

import (
	"context"
	"github.com/jokarl/go-templates/http-server/resource/example"
	"github.com/jokarl/go-templates/http-server/router"
	"github.com/jokarl/go-templates/http-server/server"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Root context listens for OS signals to gracefully shut down the server.
	rootCtx, stop := signal.NotifyContext(
		context.Background(), os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	defer stop()

	l := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	er := example.NewExampleResource(l)

	rt := router.New(
		er.Routes(),
		[]router.Middleware{},
		router.WithLogger(l),
	)

	srv := server.New(rt, server.WithLogger(l), server.WithAddr(":8080"))

	if err := srv.Start(); err != nil {
		l.Error("Failed to start server", "error", err)
	}

	<-rootCtx.Done()
	stop()
	srv.GracefulShutdown()
}
