package router

import (
	"github.com/jokarl/go-templates/http-server/logger"
	"log/slog"
	"net/http"
)

type Router struct {
	Mux        *http.ServeMux
	logger     logger.Logger
	middleware []Middleware
}

func New(opts ...func(router *Router)) *Router {
	r := &Router{
		Mux:    http.NewServeMux(),
		logger: logger.New(slog.LevelInfo),
	}

	// Apply all options
	for _, opt := range opts {
		opt(r)
	}

	return r
}

type Route struct {
	Method  Method
	Path    string
	Handler http.Handler
}

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
