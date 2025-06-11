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
	ctx             context.Context
}

const (
	defaultAddr         = ":8080"
	defaultReadTimeout  = 30 * time.Second
	defaultWriteTimeout = 30 * time.Second
)

// New creates a new server with the given options.
// The context is used for the BaseContext of http.Server
// and for graceful shutdown handling.
func New(ctx context.Context, router *router.Router, opts ...func(*Server)) *Server {
	s := &Server{
		httpServer: &http.Server{
			Addr:         defaultAddr,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
			Handler:      router.Mux,
			BaseContext: func(listener net.Listener) context.Context {
				return context.WithValue(ctx, "listener", listener)
			},
		},
		logger: logger.New(slog.LevelInfo),
		stopCh: make(chan os.Signal),
		errCh:  make(chan error),
		ctx:    ctx,
	}

	// Apply all options
	for _, opt := range opts {
		opt(s)
	}

	return s
}
