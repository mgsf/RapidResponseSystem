package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", &handler{})

	log.Fatal(http.ListenAndServe(":3000", nil))
}

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, ".."+r.URL.Path)
}
