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

//go:embed registration.html
var registrationTemplate string

var registrationPage = PageTemplate{
	Name:     "registration",
	Template: &registrationTemplate,
	Funcs:    registrationFuncMap(),
}

// Register the templates used by this handler
func init() {
	if err := RegisterTemplate(registrationPage); err != nil {
		log.Fatalf("%v template error: %v", registrationPage.Name, err)
	}
}

// Functions used by the templates
func registrationFuncMap() template.FuncMap {

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

// Registration directs the user to a page where they can sign-up or
// register to use the site
func Registration(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	// The flow is used to identify the login and registration flow and
	// return data like the csrf_token and so on.
	flow := c.QueryParam("flow")
	if flow == "" {
		c.Logger().Info("No flow ID found in URL, initializing registration flow.")
		return c.Redirect(http.StatusMovedPermanently, cc.Options.RegistrationURL())
	}

	params := public.NewGetSelfServiceRegistrationFlowParams()
	params.SetID(flow)
	res, err := cc.KratosClient().Public.GetSelfServiceRegistrationFlow(params)
	if err != nil {
		c.Logger().Error("Error getting self service registration flow, redirecting to root. Error:", err)
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
	config := res.GetPayload().Methods["password"].Config
	return c.Render(200, registrationPage.Name, map[string]interface{}{
		"config": config,
		"flow":   flow})
}
