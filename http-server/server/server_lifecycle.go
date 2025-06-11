package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Start the server.
func (s Server) Start() error {
	ctx, stop := signal.NotifyContext(s.ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.errCh <- err
		}
	}()

	s.logger.InfoContext(s.ctx, "Server started.", "address", s.httpServer.Addr)

	select {
	case err := <-s.errCh:
		close(s.errCh)
		return err
	case <-ctx.Done():
		s.logger.InfoContext(ctx, "Initiating server shutdown.", "reason", ctx.Err())

		shutdownTimeout := s.shutdownTimeout
		if shutdownTimeout == 0 {
			shutdownTimeout = s.httpServer.ReadTimeout
		}
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		s.httpServer.SetKeepAlivesEnabled(false)
		if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
			s.logger.ErrorContext(shutdownCtx, "Server shutdown error.", "error", err)
			return err
		}

		s.logger.Info("Server shutdown completed successfully.")
		return nil
	}
}
