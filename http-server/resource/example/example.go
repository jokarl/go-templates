package example

import (
	"github.com/jokarl/go-templates/http-server/router"
	"log/slog"
	"net/http"
	"path"
	"time"
)

// ExampleResource is a concrete type that implements the resource.Resource interface.
type ExampleResource struct {
	logger *slog.Logger
}

// NewExampleResource creates a new instance of ExampleResource with the provided logger.
func NewExampleResource(l *slog.Logger) *ExampleResource {
	return &ExampleResource{
		logger: l,
	}
}

func (er *ExampleResource) RootPath() string {
	return "/example"
}

func (er *ExampleResource) Routes() []router.Route {
	return []router.Route{{
		Method: router.MethodGet,
		Path:   path.Join(er.RootPath(), "/hello"),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte("Hello, World!"))
			if err != nil {
				er.logger.ErrorContext(r.Context(), "Failed to write response", "error", err)
			}
		}),
	}, {
		Method: router.MethodGet,
		Path:   path.Join(er.RootPath(), "/long"),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			select {
			case <-time.After(2 * time.Second):
				w.Write([]byte("Hello, World!"))
			case <-r.Context().Done():
				http.Error(w, "Request cancelled.", http.StatusRequestTimeout)
			}
		}),
	}}
}
