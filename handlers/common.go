package handlers

import (
	_ "embed"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

//go:embed layout.html
var layoutTemplate string

type PageTemplate struct {
	Template *string
	Name     string
	Funcs    template.FuncMap
}

// Template holds all the templates in a map
type Template struct {
	templates map[string]*template.Template
}

var TemplateRenderer = Template{
	templates: make(map[string]*template.Template),
}

// Render retrieves the templates identified by name passing data through as required
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	c.Logger().Info("render template:", name)
	if dataMap, ok := data.(map[string]interface{}); ok {
		dataMap["flash_info"] = c.QueryParam("flash_info")
		dataMap["flash_error"] = c.QueryParam("flash_error")
	}
	return t.templates[name].ExecuteTemplate(w, "layout", data)
}

// RegisterTemplate adds the template for a specific page, along with the standard layout template
func RegisterTemplate(pt PageTemplate) error {
	tmpl, err := template.New(pt.Name).Funcs(pt.Funcs).Parse(layoutTemplate)
	if err != nil {
		return err
	}
	tmpl, err = tmpl.Parse(*pt.Template)
	if err != nil {
		return err
	}
	TemplateRenderer.templates[pt.Name] = tmpl
	return err
}
