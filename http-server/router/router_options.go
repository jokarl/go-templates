package router

import "github.com/jokarl/go-templates/http-server/logger"

type Option func(router *Router)

// WithLogger sets the logger for the router.
func WithLogger(logger logger.Logger) Option {
	return func(router *Router) {
		router.logger = logger
	}
}
