package view

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/mgsf/RapidResponseSystem/config"
)

var (
	templates = map[string]*template.Template{}
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
	viewFIs, err := ioutil.ReadDir(path.Join(viewRoot, "content"))
	if err != nil {
		log.Printf("could not open view content directory: %v", err)
	}
	for _, fi := range viewFIs {
		f, err := os.Open(path.Join(viewRoot, "content", fi.Name()))
		if err != nil {
			log.Printf("failed to read content template %q: %v", fi.Name(), err)
		}
		content, err := ioutil.ReadAll(f)
		f.Close()
		if err != nil {
			log.Printf("failed to read content from template %q: %v", fi.Name(), err)
		}
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			log.Printf("failed to parse template %q: %v", fi.Name(), err)
		}
		templates[strings.TrimSuffix(fi.Name(), ".gohtml")] = tmpl
	}
}

// Get returns the view template stored with the provided key
// or an error if no template is found
func Get(key string) (*template.Template, error) {
	t, ok := templates[key]
	if !ok {
		return nil, fmt.Errorf("cannot find template with key %q", key)
	}
	return t, nil
}