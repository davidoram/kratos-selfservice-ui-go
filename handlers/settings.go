package handlers

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"

	"github.com/davidoram/kratos-selfservice-ui-go/middleware"
	"github.com/labstack/echo/v4"
	"github.com/ory/kratos-client-go/client/public"
)

//go:embed settings.html
var settingsTemplate string

var settingsPage = PageTemplate{
	Name:     "settings",
	Template: &settingsTemplate,
	Funcs:    settingsFuncMap(),
}

// Register the templates used by this handler
func init() {
	if err := RegisterTemplate(settingsPage); err != nil {
		log.Fatalf("%v template error: %v", settingsPage.Name, err)
	}
}

// Functions used by the templates
func settingsFuncMap() template.FuncMap {

	fieldLabel := map[string]string{
		"password":          "Password",
		"traits.email":      "Email",
		"traits.name.first": "First name",
		"traits.name.last":  "Last name",
	}

	return template.FuncMap{
		"labelFor": func(name string) string {
			if lbl, ok := fieldLabel[name]; ok {
				return lbl
			}
			println("No labelFor name:", name)
			return ""
		},
	}
}

// Settings directs the user to a page where they can sign-up or
// register to use the site
func Settings(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	// The flow is used to identify the login and settings flow and
	// return data like the csrf_token and so on.
	flow := c.QueryParam("flow")
	if flow == "" {
		c.Logger().Info("No flow ID found in URL, initializing settings flow.")
		return c.Redirect(http.StatusMovedPermanently, cc.Options.SettingsURL())
	}

	c.Logger().Info("Calling Kratos API to get self service settings")
	params := public.NewGetSelfServiceSettingsFlowParams()
	params.SetID(flow)

	res, err := cc.KratosAdminClient().Public.GetSelfServiceSettingsFlow(params, nil)
	if err != nil {
		c.Logger().Error("Error getting self service settings flow, redirecting to root. Error:", err)
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
	config := res.GetPayload().Methods["password"].Config
	profile := res.GetPayload().Methods["profile"].Config
	return c.Render(200, settingsPage.Name, map[string]interface{}{
		"config":  config,
		"profile": profile,
		"flow":    flow})
}
