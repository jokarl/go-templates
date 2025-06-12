package main

import (
	"context"
	"github.com/jokarl/go-templates/http-server/logger"
	"github.com/jokarl/go-templates/http-server/resource/example"
	"github.com/jokarl/go-templates/http-server/router"
	"github.com/jokarl/go-templates/http-server/server"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	sCtx, stop := signal.NotifyContext(
		ctx, os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	l := logger.New(slog.LevelInfo)

	er := example.NewExampleResource(l)

	rt := router.New(
		er.Routes(),
		[]router.Middleware{},
		router.WithLogger(l),
	)

	srv := server.New(
		sCtx,
		rt,
		server.WithLogger(l),
		server.WithAddr(":8080"),
	)

	if err := srv.Start(sCtx, stop); err != nil {
		l.Error("Server error.", "error", err)
	}

	defer os.Exit(0)
	return
}
