package router

import (
	"github.com/jokarl/go-templates/http-server/logger"
	"net/http"
	"time"
)

type Middleware func(next http.Handler) http.Handler

func loggingMiddleware(logger logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			args := []any{
				"method", r.Method,
				"path", r.URL.Path,
				"remote_addr", r.RemoteAddr,
			}

			logger.InfoContext(r.Context(), "Request received.", args...)

			next.ServeHTTP(w, r)

			duration := time.Since(startTime)
			durArgs := append(args, "duration", duration.String())

			logger.InfoContext(r.Context(), "Request completed.", durArgs...)
		})
	}
}
