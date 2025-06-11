package router

import "net/http"

type Middleware func(http.Handler) http.Handler

// WithMiddleware applies middleware to all routes.
func WithMiddleware(middleware ...Middleware) RouteCfg {
	return func(router *Router) {
		router.middleware = append(router.middleware, middleware...)
	}
}
