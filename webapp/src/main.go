package main

import (
	"log"
	"net/http"
	"time"

	"github.com/mgsf/RapidResponseSystem/view"
	"github.com/mgsf/RapidResponseSystem/wc"
)

func main() {
	view.RegisterStaticHandlers()
	routes.Register()
	log.Fatal(http.ListenAndServe(":3000", nil))
}
