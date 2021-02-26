package handlers

import (
	_ "embed"
	"html/template"
	"log"

	"github.com/davidoram/kratos-selfservice-ui-go/middleware"
	"github.com/labstack/echo/v4"
)

//go:embed dashboard.html
var dashboardTemplate string

var dashboardPage = PageTemplate{
	Name:     "dashboard",
	Template: &dashboardTemplate,
	Funcs:    dashboardFuncMap(),
}

// Register the templates used by this handler
func init() {
	if err := RegisterTemplate(dashboardPage); err != nil {
		log.Fatalf("%v template error: %v", dashboardPage.Name, err)
	}
}

// Functions used by the templates
func dashboardFuncMap() template.FuncMap {
	return template.FuncMap{}
}

// Dashboard page is accessible to logged in users only
func Dashboard(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	cc.Logger().Info(cc.KratosSession().JsonPretty())
	return c.Render(200, dashboardPage.Name, map[string]interface{}{
		"kratosSession": cc.KratosSession(),
		"headers":       cc.Request().Header,
	})
}
