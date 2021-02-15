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

	// Client for Krytos public API
	kratos *kratos.OryKratos

	// Response from the krytos 'whoami' endpoint
	kratosSession KratosSession
}

func CustomContextMiddleware(opt *options.Options) func(echo.HandlerFunc) echo.HandlerFunc {

	client := kratos.NewHTTPClientWithConfig(
		nil,
		&kratos.TransportConfig{
			Schemes:  []string{opt.KratosPublicURL.Scheme},
			Host:     opt.KratosPublicURL.Host,
			BasePath: opt.KratosPublicURL.Path})

	// CustomContext setup for all callers
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c, opt, client, KratosSession{}}
			return next(cc)
		}
	}
}

// KratosClient returns the Kratos API client
func (cc *CustomContext) KratosClient() *kratos.OryKratos {
	return cc.kratos
}

// SetKratosSession sets the whoami response
func (cc *CustomContext) SetKratosSession(payload *string) {
	cc.kratosSession.session = payload
}

// KratosSession returns the whoami response or nil
func (cc *CustomContext) KratosSession() KratosSession {
	return cc.kratosSession
}
