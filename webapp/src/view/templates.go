package view

import (
	"html/template"
	"log"
	"path"

	"github.com/mgsf/RapidResponseSystem/config"
)

// PipelineBase contains fields that are expected to be present for
// shared templates, such as _layout.html.
type PipelineBase struct {
	Title string
}

func init() {
	setupViews()
}

func setupViews() {
	log.Println("Loading view templates")
	viewRoot := config.Get().ViewRoot
	layout, err := template.ParseFiles(path.Join(viewRoot, "_layout.gohtml"))
	if err != nil {
		log.Printf("could not parse _layout.html: %v", err)
		return
	}
	_, err = layout.ParseFiles(
		path.Join(viewRoot, "_header.gohtml"),
		path.Join(viewRoot, "_footer.gohtml"),
	)
	if err != nil {
		log.Printf("could not parse _header.gohtml or _footer.gohtml: %v", err)
	}
}
