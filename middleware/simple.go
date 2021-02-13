package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ory/kratos-client-go/client/public"
)

// ProtectSimple middleware uses ORY Krato's `/sessions/whoami` endpoint to check if the user is signed in or not
func ProtectSimple(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*CustomContext)

		csrfCookie, csrfErr := cc.Request().Cookie("csrf_token")
		sessionCookie, sessionErr := cc.Request().Cookie("ory_kratos_session")

		if csrfErr == nil && sessionErr == nil {

			// Get user info from Krytos
			params := public.NewWhoamiParams().WithCookie(&csrfCookie.Value).WithAuthorization(&sessionCookie.Value)
			result, err := cc.KratosClient().Public.Whoami(params, nil)
			if err != nil {
				cc.Logger().Error("Error calling Kratos :", err)
			} else {
				cc.Logger().Info("Logged in payload:", result.Payload)
				// TODO Add to session
				return next(cc)
			}
		}

		cc.Logger().Info("Not logged in, redirect to ", cc.Options.LoginURL())
		return c.Redirect(http.StatusMovedPermanently, cc.Options.LoginURL())
	}
}
