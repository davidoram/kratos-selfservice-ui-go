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

//go:embed login.html
var loginTemplate string

var loginPage = PageTemplate{
	Name:     "login",
	Template: &loginTemplate,
	Funcs:    loginFuncMap(),
}

// Register the templates used by this handler
func init() {
	if err := AddPageTemplate(loginPage); err != nil {
		log.Fatalf("%v template error: %v", loginPage.Name, err)
	}
}

// Functions used by the templates
func loginFuncMap() template.FuncMap {

	fieldLabel := map[string]string{
		"password":   "Password",
		"identifier": "Email",
	}

	return template.FuncMap{
		"labelFor": func(name string) string {
			if lbl, ok := fieldLabel[name]; ok {
				return lbl
			}
			println("labelFor name:", name)
			return ""
		},
	}
}

// Login handler displays the login screen
func Login(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	// The flow is used to identify the login and registration flow and
	// return data like the csrf_token and so on.
	flow := c.QueryParam("flow")
	if flow == "" {
		c.Logger().Info("'No flow ID found in URL, initializing login flow.")
		return c.Redirect(http.StatusMovedPermanently, cc.Options.LoginFlowURL())
	}

	params := public.NewGetSelfServiceLoginFlowParams()
	params.SetID(flow)
	res, err := cc.KratosClient().Public.GetSelfServiceLoginFlow(params)
	if err != nil {
		c.Logger().Info("Login error=", err)
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
	config := res.GetPayload().Methods["password"].Config
	c.Logger().Info("Login with config:", config)
	return c.Render(200, loginPage.Name, map[string]interface{}{
		"config": config,
		"flow":   flow})

}
