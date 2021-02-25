package main

import (
	"log"

	"github.com/davidoram/kratos-selfservice-ui-go/handlers"
	"github.com/davidoram/kratos-selfservice-ui-go/middleware"
	"github.com/davidoram/kratos-selfservice-ui-go/options"
	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"
	elog "github.com/labstack/gommon/log"
)

func main() {
	//log.Printf("env: %v", os.Environ())
	opt := options.NewOptions().SetFromCommandLine()
	if err := opt.Validate(); err != nil {
		log.Fatalf("Error parsing command line: %v", err)
	}
	log.Printf("KratosAdminURL: %s", opt.KratosAdminURL.String())
	log.Printf("KratosPublicURL: %s", opt.KratosPublicURL.String())
	log.Printf("KratosBrowserURL: %s", opt.KratosPublicURL.String())
	log.Printf("BaseURL: %s", opt.BaseURL.String())
	log.Printf("Port: %d", opt.Port)

	// Echo instance
	e := echo.New()
	e.HideBanner = true

	// Configure logging
	e.Logger.SetLevel(elog.INFO)
	e.Logger.SetHeader("${time_rfc3339} ${short_file}:${line}")

	// Custom renderer
	e.Renderer = &handlers.TemplateRenderer

	// Common middleware
	e.Use(emiddleware.RequestID(),
		emiddleware.Recover(),
		emiddleware.LoggerWithConfig(emiddleware.LoggerConfig{
			Format: "${time_rfc3339} ${id} ${method} ${path} - ${msg} \n",
		}),
		middleware.CustomContextMiddleware(opt))

	// Routes
	e.GET("/", handlers.Home, middleware.NoCache())
	e.GET("/dashboard", handlers.Dashboard, middleware.NoCache(), middleware.ProtectSimple)
	e.GET("/auth/registration", handlers.Registration, middleware.NoCache())
	e.GET("/auth/settings", handlers.Settings, middleware.NoCache())
	e.GET("/auth/login", handlers.Login, middleware.NoCache())
	e.GET("/auth/logout", handlers.Logout, middleware.NoCache())

	// Start server
	e.Logger.Fatal(e.Start(opt.Address()))

}
