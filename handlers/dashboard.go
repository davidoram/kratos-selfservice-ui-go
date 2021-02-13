package handlers

import (
	"net/http"

	"github.com/davidoram/kratos-selfservice-ui-go/middleware"
	"github.com/labstack/echo/v4"
)

// Dashboard page is accessible to logged in users only
func Dashboard(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	// The flow is used to identify the login and registration flow and
	// return data like the csrf_token and so on.
	flow := c.QueryParam("flow")
	if flow == "" {
		return c.Redirect(http.StatusMovedPermanently, cc.Options.RegistrationURL())
	}

	return c.String(http.StatusOK, "Dashboard")
}
