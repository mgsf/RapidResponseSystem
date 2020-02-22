package view

import (
	"net/http"

	"github.com/mgsf/RapidResponseSystem/config"
)

// RegisterStaticHandlers registers HTTP handlers that will serve static
// content such as CSS and JavaScript files
func RegisterStaticHandlers() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(config.Get().StaticRoot))))
}
