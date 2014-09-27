package routing

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

// resourceFn defines a function type that is used to register a namespaced set
// of resources
type ResourceFn func(string) *web.Mux

// add a resource (set of endpoints) to goji router
func AddResource(root string, fn ResourceFn) {
	goji.Handle(fmt.Sprintf("%s/*", root), fn(root))
	goji.Get(root, http.RedirectHandler(fmt.Sprintf("%s/", root), 301))
}

// get the fully namespaced path for a request path using the specified prefix
func ResourcePath(prefix, path string) string {
	return fmt.Sprintf("%s%s", prefix, path)
}
