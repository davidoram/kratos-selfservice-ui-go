package handlers

import (
	"net/http"

	"github.com/davidoram/kratos-selfservice-ui-go/middleware"
	"github.com/labstack/echo/v4"
)

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

	return c.String(http.StatusOK, "Login todo")
}
