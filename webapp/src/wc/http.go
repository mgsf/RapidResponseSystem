package wc

import (
	"log"
	"net/http"
	"regexp"

	"github.com/silvercloudtraining/webapp/view"
)

type httpHandler struct {
	urlPattern *regexp.Regexp
}

// NewViewHandler returns a handler that handles requests for HTML files
// related to workcenters.
func NewViewHandler() http.Handler {
	return &httpHandler{
		urlPattern: regexp.MustCompile(`^\/wc\/(\d+)$`),
	}
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// hardcode id for now
	id := 1

	t, err := view.Get("workcenter")
	if err != nil {
		log.Println("unable to find view template for workcenter")
		http.Error(w, "view template not found", http.StatusInternalServerError)
		return
	}
	wc, err := GetWorkcenter(id)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	w.Header().Add("Content-Type", "text/html")
	err = t.Execute(w, struct {
		Workcenter
		view.PipelineBase
	}{
		Workcenter:   wc,
		PipelineBase: view.PipelineBase{Title: wc.Name},
	})
	if err != nil {
		log.Print(err)
		http.Error(w, "failed to generate view", http.StatusInternalServerError)
	}
}
