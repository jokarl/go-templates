package server

import (
	"context"
	"github.com/jokarl/go-templates/http-server/logger"
	"github.com/jokarl/go-templates/http-server/router"
	"log/slog"
	"net"
	"net/http"
	"time"
)

// Server represents a simple HTTP server with graceful shutdown capabilities.
type Server struct {
	httpServer       *http.Server
	logger           *slog.Logger
	shutdownTimeout  time.Duration
	gracefulShutdown context.CancelFunc
}

const (
	defaultAddr         = ":8080"
	defaultReadTimeout  = 30 * time.Second
	defaultWriteTimeout = 30 * time.Second
)

// New creates a new server with the given options.
func New(router *router.Router, opts ...func(*Server)) *Server {
	// Ensure in-flight requests aren't cancelled immediately on SIGTERM
	ongoingCtx, stopOngoingGracefully := context.WithCancel(context.Background())
	s := &Server{
		httpServer: &http.Server{
			Addr:         defaultAddr,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
			Handler:      router,
			BaseContext: func(listener net.Listener) context.Context {
				return ongoingCtx
			},
		},
		logger:           logger.New(slog.LevelInfo),
		gracefulShutdown: stopOngoingGracefully,
	}

	// Apply all options
	for _, opt := range opts {
		opt(s)
	}

	return s
}
