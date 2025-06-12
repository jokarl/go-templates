package resource

import "github.com/jokarl/go-templates/http-server/router"

// Resource defines the interface for a resource in the HTTP server.
// A resource is expected to provide a root path and a list of routes.
type Resource interface {
	RootPath() string
	Routes() []router.Route
}
