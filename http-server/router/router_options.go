package router

import (
	"log/slog"
)

type Option func(router *Router)

// WithLogger sets the logger for the router.
func WithLogger(logger *slog.Logger) Option {
	return func(router *Router) {
		router.logger = logger
	}
}
