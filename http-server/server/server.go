package server

import (
	"github.com/jokarl/go-templates/http-server/logger"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// Server represents a simple HTTP server with graceful shutdown capabilities.
type Server struct {
	httpServer *http.Server
	logger     logger.Logger
	stopCh     chan os.Signal
	errCh      chan error
}

const (
	defaultAddr         = ":8080"
	defaultReadTimeout  = 30 * time.Second
	defaultWriteTimeout = 30 * time.Second
)

// New creates a new server with the given options.
func New(opts ...func(*Server)) *Server {
	s := &Server{
		httpServer: &http.Server{
			Addr:         defaultAddr,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
			Handler:      http.NewServeMux(),
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
