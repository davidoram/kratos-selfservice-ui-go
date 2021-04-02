package handlers

import (
	"bytes"
	_ "embed"
	"html/template"
	"io"
	"log"
	"net/http"
)

// Template wraps an html/template
type Template struct {
	tmpl *template.Template
}

var (
	templateMap = make(map[TemplateName]Template)
)

const (
	// ErrRenderingPage is the text displayed if we get an error during rendering
	ErrRenderingPage = "Error rendering page"
)

// RegisterTemplate creates a html template with a name, and a FuncMap for a set of template strings,
// along with a list of 'templates' which will include the 'layout', 'header', 'footer' and 'content' etc
func RegisterTemplate(name TemplateName, fmap template.FuncMap, templates ...string) error {
	var err error
	tmpl := template.New(string(name)).Funcs(fmap)
	for _, t := range templates {
		tmpl, err = tmpl.Parse(t)
		if err != nil {
			return err
		}
	}
	templateMap[name] = Template{tmpl: tmpl}
	return err
}

// GetTemplate returns template with name or nil
func GetTemplate(name TemplateName) Template {
	return templateMap[name]
}

// Render executes the template 'name' passing dataMap
func (t Template) Render(name string, w http.ResponseWriter, r *http.Request, dataMap map[string]interface{}) error {
	log.Printf("Render template: %s", t.tmpl.Name())

	// Add common query params into the dataMap
	dataMap["flash_info"] = r.URL.Query().Get("flash_info")
	dataMap["flash_error"] = r.URL.Query().Get("flash_error")

	// Render to a buffer
	var b bytes.Buffer
	err := t.tmpl.ExecuteTemplate(&b, name, dataMap)
	if err != nil {
		log.Printf("Error executing template: %s\n", err)
		http.Error(w, ErrRenderingPage, http.StatusInternalServerError)
		return err
	}

	// Copy the buffer to the HTML writer
	size, err := io.Copy(w, &b)
	if err != nil {
		log.Printf("Error copying template: %s, bytes %d\n", err, size)
		http.Error(w, ErrRenderingPage, http.StatusInternalServerError)
		return err
	}
	return nil
}
