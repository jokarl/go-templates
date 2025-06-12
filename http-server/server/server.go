package server

import (
	"context"
	"github.com/jokarl/go-templates/http-server/logger"
	"github.com/jokarl/go-templates/http-server/router"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
)

// Server represents a simple HTTP server with graceful shutdown capabilities.
type Server struct {
	httpServer      *http.Server
	logger          logger.Logger
	stopCh          chan os.Signal
	errCh           chan error
	shutdownTimeout time.Duration
}

const (
	defaultAddr         = ":8080"
	defaultReadTimeout  = 30 * time.Second
	defaultWriteTimeout = 30 * time.Second
)

// New creates a new server with the given options.
func New(ctx context.Context, router *router.Router, opts ...func(*Server)) *Server {
	s := &Server{
		httpServer: &http.Server{
			Addr:         defaultAddr,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
			Handler:      router,
			BaseContext: func(listener net.Listener) context.Context {
				return context.WithValue(ctx, "listener", listener)
			},
		},
		logger: logger.New(slog.LevelInfo),
		stopCh: make(chan os.Signal),
		errCh:  make(chan error),
	}

	// Apply all options
	for _, opt := range opts {
		opt(s)
	}

	return s
}
