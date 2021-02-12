package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Home displays a simple homepage
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Homepage")
}
