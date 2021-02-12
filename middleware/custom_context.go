package middleware

import (
	"github.com/davidoram/kratos-selfservice-ui-go/options"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Options options.Options
}
