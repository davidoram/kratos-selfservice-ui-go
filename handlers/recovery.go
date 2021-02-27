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

//go:embed recovery.html
var recoveryTemplate string

var recoveryPage = PageTemplate{
	Name:     "recovery",
	Template: &recoveryTemplate,
	Funcs:    recoveryFuncMap(),
}

// Register the templates used by this handler
func init() {
	if err := RegisterTemplate(recoveryPage); err != nil {
		log.Fatalf("%v template error: %v", recoveryPage.Name, err)
	}
}

// Functions used by the templates
func recoveryFuncMap() template.FuncMap {

	fieldLabel := map[string]string{
		"email": "Email",
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

// Recovery handler displays the recovery screen, which allows the user to enter
// their email address, an email will be sent there that will automatically authenticate then
func Recovery(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	// The flow is used to identify the recovery and registration flow and
	// return data like the csrf_token and so on.
	flow := c.QueryParam("flow")
	if flow == "" {
		c.Logger().Info("'No flow ID found in URL, initializing recovery flow.")
		return c.Redirect(http.StatusMovedPermanently, cc.Options.RecoveryFlowURL())
	}

	params := public.NewGetSelfServiceRecoveryFlowParams()
	params.SetID(flow)
	c.Logger().Info("Calling Kratos API to get self service recovery")
	res, err := cc.KratosPublicClient().Public.GetSelfServiceRecoveryFlow(params)
	if err != nil {
		c.Logger().Error("Error getting self service recovery flow, redirecting to root. Error:", err)
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
	link := res.GetPayload().Methods["link"].Config
	state := res.GetPayload().State
	c.Logger().Info("state=", state)
	return c.Render(200, recoveryPage.Name, map[string]interface{}{
		"link":  link,
		"flow":  flow,
		"state": state})

}
