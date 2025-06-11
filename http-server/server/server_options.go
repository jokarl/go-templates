package server

import (
	"github.com/jokarl/go-templates/http-server/logger"
	"time"
)

// Option is a function that configures the server
type Option func(*Server)

// WithLogger sets the logger to use for the server
func WithLogger(logger logger.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

// WithAddr sets the server address
func WithAddr(addr string) Option {
	return func(s *Server) {
		s.httpServer.Addr = addr
	}
}

// WithReadTimeout sets the read timeout.
func WithReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.httpServer.ReadTimeout = timeout
	}
}

// WithShutdownTimeout sets the timeout for graceful shutdown.
// If not set, defaults to WithReadTimeout.
func WithShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}

// WithWriteTimeout sets the write timeout.
func WithWriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.httpServer.WriteTimeout = timeout
	}
}
