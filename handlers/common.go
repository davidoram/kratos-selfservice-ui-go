package handlers

import (
	_ "embed"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Template holds all the templates in a map
type Template struct {
	templates map[string]*template.Template
}

var t = Template{
	templates: make(map[string]*template.Template),
}

// Render renders templates passing data through as required
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates[name].ExecuteTemplate(w, name, data)
}

// InitTemplates initialises each template
func InitTemplates() error {
	var err error
	t.templates["login"], err = template.New("login").Parse(loginTemplate)
	return err
}
