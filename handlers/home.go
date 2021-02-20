package handlers

import (
	_ "embed"
	"html/template"
	"log"

	"github.com/davidoram/kratos-selfservice-ui-go/middleware"
	"github.com/labstack/echo/v4"
)

//go:embed home.html
var homeTemplate string

var homePage = PageTemplate{
	Name:     "home",
	Template: &homeTemplate,
	Funcs:    homeFuncMap(),
}

// Register the templates used by this handler
func init() {
	if err := RegisterTemplate(homePage); err != nil {
		log.Fatalf("%v template error: %v", homePage.Name, err)
	}
}

// Functions used by the templates
func homeFuncMap() template.FuncMap {
	return template.FuncMap{}
}

// Home displays a simple homepage
func Home(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	return c.Render(200, homePage.Name, map[string]interface{}{
		"kratosSession": cc.KratosSession(),
		"headers":       []string{},
	})
}
