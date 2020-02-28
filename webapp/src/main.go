package main

import (
	"log"
	"net/http"

	"github.com/mgsf/RapidResponseSystem/routes"
	"github.com/mgsf/RapidResponseSystem/view"
)

func main() {
	view.RegisterStaticHandlers()
	routes.Register()
	log.Fatal(http.ListenAndServe(":3000", nil))
}
