package middleware

import (
	"github.com/labstack/echo/v4"
)

// SimpleLog middleware logs the http request
func SimpleLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Logger().Info(c.Request().Method, c.Request().RequestURI)
		return next(c)
	}
}
