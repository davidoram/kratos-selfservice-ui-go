package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var homepage = `
<html>
<body>
<h1>Homepage</h1>
<a href="/dashboard">Dashboard</a><br>
<a href="/auth/registration">Registration</a><br>
<a href="/auth/login">Login</a><br>
<a href="/auth/logout">Logout</a><br>
</body>
</html>
`

// Home displays a simple homepage
func Home(c echo.Context) error {
	return c.HTML(http.StatusOK, homepage)
}
