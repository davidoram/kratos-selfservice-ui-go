package handlers

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"

	"github.com/davidoram/kratos-selfservice-ui-go/middleware"
	"github.com/labstack/echo/v4"
)

//go:embed logout.html
var logoutTemplate string

var logoutPage = PageTemplate{
	Name:     "logout",
	Template: &logoutTemplate,
	Funcs:    logoutFuncMap(),
}

// Register the templates used by this handler
func init() {
	if err := AddPageTemplate(logoutPage); err != nil {
		log.Fatalf("%v template error: %v", logoutPage.Name, err)
	}
}

// Functions used by the templates
func logoutFuncMap() template.FuncMap {
	return template.FuncMap{}
}

// Logout handler displays the logout screen
func Logout(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	// The flow is used to identify the logout and registration flow and
	// return data like the csrf_token and so on.
	flow := c.QueryParam("flow")
	if flow == "" {
		c.Logger().Info("'No flow ID found in URL, initializing logout flow.")
		return c.Redirect(http.StatusMovedPermanently, cc.Options.LogoutFlowURL())
	}

	return c.Render(200, logoutPage.Name, map[string]interface{}{})

}
