package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/davidoram/kratos-selfservice-ui-go/handlers"
	"github.com/davidoram/kratos-selfservice-ui-go/middleware"
	"github.com/davidoram/kratos-selfservice-ui-go/options"

	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func main() {

	opt := options.NewOptions().SetFromCommandLine()
	if err := opt.Validate(); err != nil {
		log.Fatalf("Error parsing command line: %v", err)
	}
	log.Printf("KratosAdminURL: %s", opt.KratosAdminURL.String())
	log.Printf("KratosPublicURL: %s", opt.KratosPublicURL.String())
	log.Printf("KratosBrowserURL: %s", opt.KratosPublicURL.String())
	log.Printf("BaseURL: %s", opt.BaseURL.String())
	log.Printf("Port: %d", opt.Port)

	// Setup sesssion store in cookies
	var store = sessions.NewCookieStore([]byte("'+VO7Qir8yZ9idvyPktSBbmaVjDvNw9fRlXTdIBO9FqI=")) // TODO make opt for session key

	// Public Routes need no authentication
	//
	r := mux.NewRouter()
	r.Use(gh.RecoveryHandler(),
		// gh.LoggingHandler().ServeHTTP,
		middleware.NoCacheMiddleware)

	r.HandleFunc("/", handlers.Home)
	regP := handlers.RegistrationParams{
		FlowRedirectURL: opt.RegistrationURL(),
	}
	r.HandleFunc("/auth/registration", regP.Registration)
	settingsP := handlers.SettingsParams{
		FlowRedirectURL: opt.SettingsURL(),
	}
	r.HandleFunc("/auth/settings", settingsP.Settings)
	loginP := handlers.LoginParams{
		FlowRedirectURL: opt.LoginFlowURL(),
	}
	r.HandleFunc("/auth/login", loginP.Login)
	logoutP := handlers.LogoutParams{
		FlowRedirectURL: opt.LogoutFlowURL(),
	}
	r.HandleFunc("/auth/logout", logoutP.Logout)
	recoverP := handlers.RecoveryParams{
		FlowRedirectURL: opt.RecoveryFlowURL(),
	}
	r.HandleFunc("/auth/recovery", recoverP.Recovery)

	// Following routes must be authenticated, so they get extra middleware
	//
	authP := middleware.KratosAuthParams{
		WhoAmIURL:         opt.WhoAmIURL(),
		RedirectUnauthURL: "/",
		Store:             store,
	}
	dRoute := r.HandleFunc("/dashboard", handlers.Dashboard)
	dRoute.Subrouter().Use(authP.KratoAuthMiddleware)

	// Start server
	srv := &http.Server{
		Addr: fmt.Sprintf("localhost:%d", opt.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), opt.ShutdownWait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
