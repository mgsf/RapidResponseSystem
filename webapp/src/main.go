package main

import (
	"log"
	"net/http"

	"github.com/mgsf/RapidResponseSystem/view"
)

func main() {
	view.RegisterStaticHandlers()
	log.Fatal(http.ListenAndServe(":3000", nil))
}
