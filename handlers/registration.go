package handlers

import (
	"net/http"

	"github.com/davidoram/kratos-selfservice-ui-go/middleware"
	"github.com/labstack/echo/v4"
)

// Registration directs the user to a page where they can sign-up or
// register to use the site
func Registration(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	// The flow is used to identify the login and registration flow and
	// return data like the csrf_token and so on.
	flow := c.QueryParam("flow")
	if flow == "" {
		return c.Redirect(http.StatusMovedPermanently, cc.Options.RegistrationURL())
	}

	return c.String(http.StatusOK, "Registration todo")
}
