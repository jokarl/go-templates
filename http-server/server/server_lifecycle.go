package server

import (
	"context"
	"errors"
	"net/http"
	"sync/atomic"
	"time"
)

var isShuttingDown atomic.Bool

const (
	_shutdownPeriod      = 15 * time.Second
	_shutdownHardPeriod  = 3 * time.Second
	_readinessDrainDelay = 5 * time.Second
)

// Start the server.
func (s Server) Start() {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
}

// GracefulShutdown initiates a graceful shutdown of the server.
//
// It allows _readinessDrainDelay to propagate readiness checks.
// After that, it waits for ongoing requests to finish within _shutdownPeriod.
// If the shutdown period expires, it waits for an additional _shutdownHardPeriod
func (s Server) GracefulShutdown() {
	isShuttingDown.Store(true)
	s.logger.Info("Received shutdown signal, shutting down.")

	// Give time for readiness check to propagate
	time.Sleep(_readinessDrainDelay)
	s.logger.Info("Readiness check propagated, now waiting for ongoing requests to finish.")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), _shutdownPeriod)
	defer cancel()
	err := s.httpServer.Shutdown(shutdownCtx)
	s.gracefulShutdown()
	if err != nil {
		s.logger.Error("Failed to wait for ongoing requests to finish, waiting for forced cancellation.")
		time.Sleep(_shutdownHardPeriod)
	}

	s.logger.Info("Server shut down gracefully.")
}
