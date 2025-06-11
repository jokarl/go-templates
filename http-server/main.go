package main

import (
	"github.com/jokarl/go-templates/http-server/logger"
	"github.com/jokarl/go-templates/http-server/router"
	"github.com/jokarl/go-templates/http-server/server"
	"log/slog"
	"net/http"
)

func main() {
	l := logger.New(slog.LevelInfo)

	r := router.New(
		router.WithLogger(l),
		router.WithRoutes(router.Route{
			Method: router.MethodGet,
			Path:   "/hello",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, err := w.Write([]byte("Hello, World!"))
				if err != nil {
					l.Error("Failed to write response", "error", err)
				}
			}),
		}),
	)

	srv := server.New(
		server.WithLogger(l),
		server.WithAddr(":8080"),
		server.WithRouter(r),
	)

	if err := srv.Start(); err != nil {
		l.Error("Server error.", "error", err)
	}
}
