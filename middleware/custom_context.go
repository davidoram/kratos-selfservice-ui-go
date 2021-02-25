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
	kratosPublic *kratos.OryKratos

	// Response from the krytos 'whoami' endpoint
	kratosSession KratosSession

	// Client for Krytos admin API
	kratosAdmin *kratos.OryKratos
}

func CustomContextMiddleware(opt *options.Options) func(echo.HandlerFunc) echo.HandlerFunc {

	publicClient := kratos.NewHTTPClientWithConfig(
		nil,
		&kratos.TransportConfig{
			Schemes:  []string{opt.KratosPublicURL.Scheme},
			Host:     opt.KratosPublicURL.Host,
			BasePath: opt.KratosPublicURL.Path})

	adminClient := kratos.NewHTTPClientWithConfig(
		nil,
		&kratos.TransportConfig{
			Schemes:  []string{opt.KratosAdminURL.Scheme},
			Host:     opt.KratosAdminURL.Host,
			BasePath: opt.KratosAdminURL.Path})

	// CustomContext setup for all callers
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c, opt, publicClient, KratosSession{}, adminClient}
			return next(cc)
		}
	}
}

// KratosPublicClient returns the Kratos Public API client
func (cc *CustomContext) KratosPublicClient() *kratos.OryKratos {
	return cc.kratosPublic
}

// SetKratosSession sets the whoami response
func (cc *CustomContext) SetKratosSession(payload *string) {
	cc.kratosSession.session = payload
}

// KratosSession returns the whoami response or nil
func (cc *CustomContext) KratosSession() KratosSession {
	return cc.kratosSession
}

// KratosAdminClient returns the Kratos Admin API client
func (cc *CustomContext) KratosAdminClient() *kratos.OryKratos {
	return cc.kratosAdmin
}
