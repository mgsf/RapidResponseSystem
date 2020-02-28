package routes

import (
	"net/http"

	"github.com/silvercloudtraining/webapp/wc"
)

// Register connects all of the Handlers in the application
// to their respective routes. Parametric and child routes
// will be implemented by these base handlers.
func Register() {
	http.Handle("/", http.NotFoundHandler())
	http.Handle("/wc", wc.NewViewHandler())
}
