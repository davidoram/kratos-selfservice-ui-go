package main

import (
	"log"

	"github.com/davidoram/kratos-selfservice-ui-go/handlers"
	"github.com/davidoram/kratos-selfservice-ui-go/options"
	"github.com/labstack/echo/v4"
)

func main() {
	opt := options.NewOptions().SetFromCommandLine()
	if err := opt.Validate(); err != nil {
		log.Fatalf("Error parsing command line: %v", err)
	}
	log.Printf("KratosAdminURL: %s", opt.KratosAdminURL.String())
	log.Printf("KratosPublicURL: %s", opt.KratosPublicURL.String())
	log.Printf("BaseURL: %s", opt.BaseURL.String())
	log.Printf("Port: %d", opt.Port)

	// Echo instance
	e := echo.New()
	e.HideBanner = true

	// Routes
	e.GET("/", handlers.Home)
	e.GET("/auth/registration", handlers.Registration)

	// Start server
	e.Logger.Fatal(e.Start(opt.Address()))

}
