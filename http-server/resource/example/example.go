package example

import (
	"github.com/jokarl/go-templates/http-server/logger"
	"github.com/jokarl/go-templates/http-server/router"
	"net/http"
	"path"
)

// ExampleResource is a concrete type that implements the Resource interface.
type ExampleResource struct {
	logger logger.Logger
}

// NewExampleResource creates a new instance of ExampleResource with the provided logger.
func NewExampleResource(l logger.Logger) *ExampleResource {
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
	}}
}
