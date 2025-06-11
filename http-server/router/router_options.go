package router

import "github.com/jokarl/go-templates/http-server/logger"

type RouteCfg func(router *Router)

// WithRoutes registers multiple routes with the router.
func WithRoutes(routes ...Route) RouteCfg {
	return func(router *Router) {
		for _, route := range routes {
			router.Mux.Handle(string(route.Method)+" "+route.Path, route.Handler)
			router.logger.Info("Route registered.", "method", route.Method, "path", route.Path)
		}
	}
}

// WithLogger sets the logger for the router.
func WithLogger(logger logger.Logger) RouteCfg {
	return func(router *Router) {
		router.logger = logger
	}
}
