package router

import (
	"github.com/jokarl/go-templates/http-server/logger"
	"log/slog"
	"net/http"
)

// Router represents an HTTP router that manages routes and middlewares.
type Router struct {
	mux         *http.ServeMux
	logger      logger.Logger
	middlewares []Middleware
}

// New creates a new Router instance with the provided middlewares and options.
// The first middleware is always a logging middleware that logs requests and responses.
func New(routes []Route, mw []Middleware, opts ...func(router *Router)) *Router {
	rt := &Router{
		mux:         http.NewServeMux(),
		logger:      logger.New(slog.LevelInfo),
		middlewares: mw,
	}

	for _, opt := range opts {
		opt(rt)
	}

	rt.middlewares = append([]Middleware{loggingMiddleware(rt.logger)}, rt.middlewares...)
	rt.registerRoutes(routes)

	return rt
}

func (rt *Router) registerRoutes(routes []Route) {
	for _, route := range routes {
		h := route.Handler

		for i := len(rt.middlewares) - 1; i >= 0; i-- {
			if rt.middlewares[i] != nil {
				h = rt.middlewares[i](h)
			}
		}

		rt.mux.Handle(string(route.Method)+" "+route.Path, h)
		rt.logger.Info("Route registered.",
			"method", route.Method,
			"path", route.Path)
	}
}

// ServeHTTP makes the Router an http.Handler.
func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt.mux.ServeHTTP(w, r)
}

// Route defines a single HTTP route.
type Route struct {
	Method  Method
	Path    string
	Handler http.Handler
}

// Method represents an HTTP method.
type Method string

const (
	MethodGet     Method = "GET"
	MethodHead    Method = "HEAD"
	MethodPost    Method = "POST"
	MethodPut     Method = "PUT"
	MethodPatch   Method = "PATCH"
	MethodDelete  Method = "DELETE"
	MethodConnect Method = "CONNECT"
	MethodOptions Method = "OPTIONS"
	MethodTrace   Method = "TRACE"
)
