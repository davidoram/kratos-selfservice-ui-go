package handlers

import (
	_ "embed"
	"net/http"

	"github.com/davidoram/kratos-selfservice-ui-go/middleware"
	"github.com/labstack/echo/v4"
	"github.com/ory/kratos-client-go/client/public"
)

//go:embed login.html
var loginTemplate string

// Login shows the login / registration screen
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
	c.Logger().Info("login ", res)

	// return c.Render(200, "auth_login/layout.html", map[string]interface{}{
	// 	"kratosSession":   c.kratosSession,
	// 	"MetaDescription": "Planting Wizard",
	// 	"PageTitle":       "Planting Wizard Verify",
	// 	"NavCurrent":      "login",
	// 	"KratosConfig":    res.GetPayload().Methods["password"].Config,
	// 	"flow":            flow,
	// 	"kratosWebUrl":    c.KratosWebUrl(),
	// 	"csrf":            csrfToken(c),
	// })

	return c.String(http.StatusOK, "Login todo")
}
