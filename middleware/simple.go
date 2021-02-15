package middleware

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ProtectSimple middleware uses ORY Krato's `/sessions/whoami` endpoint to check if the user is signed in or not
func ProtectSimple(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*CustomContext)
		cc.Logger().Info("ProtectSimple middleware")

		csrfCookie, csrfErr := cc.Request().Cookie("csrf_token")
		sessionCookie, sessionErr := cc.Request().Cookie("ory_kratos_session")

		if csrfErr == nil && sessionErr == nil {

			client := &http.Client{}
			if req, err := http.NewRequest("POST", cc.Options.WhoAmIURL(), nil); err != nil {
				c.Logger().Error("Error creating request, error:", err)
			} else {
				// Add Cookies
				req.AddCookie(csrfCookie)
				req.AddCookie(sessionCookie)

				// Setup Headers
				req.Header.Add("Accept", "application/json")

				// Make the API call
				if res, err := client.Do(req); err != nil {
					c.Logger().Error("API returned error:", err)
				} else {
					defer res.Body.Close()
					if body, err := ioutil.ReadAll(res.Body); err != nil {
						c.Logger().Error("ReadAll error:", err)
					} else {
						if res.StatusCode != 200 {
							c.Logger().Error("Status code not 200, status_code", res.StatusCode)
						} else {
							s := string(body)
							cc.SetKratosSession(&s)
							cc.Logger().Info("Logged in to kratos")
							return next(cc)
						}
					}
				}
			}
		}

		cc.Logger().Info("Not logged in, redirect to ", cc.Options.LoginPageURL())
		return c.Redirect(http.StatusMovedPermanently, cc.Options.LoginPageURL())
	}
}
