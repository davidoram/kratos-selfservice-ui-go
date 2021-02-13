package handlers

import (
	_ "embed"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

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

// Render renders templates passing data through as required
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	c.Logger().Info("render template:", name)
	return t.templates[name].ExecuteTemplate(w, name, data)
}

// AddPageTemplate adds the template for a specific page
func AddPageTemplate(pt PageTemplate) error {
	tmpl, err := template.New(pt.Name).Funcs(pt.Funcs).Parse(*pt.Template)
	if err != nil {
		return err
	}
	TemplateRenderer.templates[pt.Name] = tmpl
	return err
}

// func registrationFuncMap() template.FuncMap {

// 	fieldLabel := map[string]string{
// 		"password":          "Password",
// 		"traits.email":      "Email",
// 		"traits.name.first": "First name",
// 		"traits.name.last":  "Surname",
// 	}

// 	return template.FuncMap{
// 		"labelFor": func(name string) string {
// 			if lbl, ok := fieldLabel[name]; ok {
// 				return lbl
// 			}
// 			return ""
// 		},
// 	}
// }
