package middleware

import (
	"github.com/davidoram/kratos-selfservice-ui-go/options"
	"github.com/labstack/echo/v4"
	kratos "github.com/ory/kratos-client-go/client"
)

type CustomContext struct {
	echo.Context

	// Options holds the runtime configuration
	Options *options.Options

	// kratos is a client used to call Krytos public API
	kratos *kratos.OryKratos
}

func CustomContextMiddleware(opt *options.Options) func(echo.HandlerFunc) echo.HandlerFunc {

	client := kratos.NewHTTPClientWithConfig(
		nil,
		&kratos.TransportConfig{
			Schemes:  []string{opt.KratosAdminURL.Scheme},
			Host:     opt.KratosAdminURL.Host,
			BasePath: opt.KratosAdminURL.Path})

	// CustomContext setup for all callers
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c, opt, client}
			return next(cc)
		}
	}
}

// KratosClient returns the Kratos API client
func (cc *CustomContext) KratosClient() *kratos.OryKratos {
	return cc.kratos
}
