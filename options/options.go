package options

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/securecookie"
)

// Options holds the application command line options
type Options struct {

	// KratosAdminURL is the URL where ORY Kratos's Admin API is located at.
	// If this app and ORY Kratos are running in the same private network, this should be the
	// private network address (e.g. kratos-admin.svc.cluster.local).
	KratosAdminURL *url.URL

	// KratosPublicURL is the URL where ORY Kratos's Public API is located at.
	// If this app and ORY Kratos are running in the same private network, this should be the
	// private network address (e.g. kratos-public.svc.cluster.local).
	KratosPublicURL *url.URL

	// KratosBrowserURL is the URL where ORY Kratos's self service browser endpoints are located at.
	KratosBrowserURL *url.URL

	// BaseURL is the base url of this app. If served e.g. behind a proxy or via GitHub pages
	// this would be the path, e.g. https://mywebsite.com/kratos-selfservice-ui-go/. Must be absolute!
	BaseURL *url.URL

	// Host that the app is listening on. Used together with Port
	Host string

	// Port that this app is listening on. Used together with Host
	Port int

	// Duration to wait when asked to shutdown gracefully
	ShutdownWait time.Duration

	// TLSCertPath is an optional Path to certificate file.
	// Should be set up together with TLSKeyPath to enable HTTPS.
	TLSCertPath string
	// TLSCertPath is an optional path to key file.
	// Should be set up together with TLSCertPath to enable HTTPS.
	TLSKeyPath string

	// Pairs of authentication and encryption keys for Cookies
	CookieStoreKeyPairs [][]byte
}

func NewOptions() *Options {
	return &Options{
		KratosAdminURL:   &url.URL{},
		KratosPublicURL:  &url.URL{},
		KratosBrowserURL: &url.URL{},
		BaseURL:          &url.URL{},
	}
}

// SetFromCommandLine will parse the command line, and populate the Options.
// The special case is when the 'gen-cookie-store-key-pair' is detected, will genrate the keys and exit
// Will also exit if key-pairs passed in are invalid
func (o *Options) SetFromCommandLine() *Options {

	KratosAdminURL := MustMakeURLValue(os.Getenv("KRATOS_ADMIN_URL"))
	flag.Var(&KratosAdminURL, "kratos-admin-url", "The URL where ORY Kratos's Admin API is located at. If this app and ORY Kratos are running in the same private network, this should be the private network address. Defaults to KRATOS_ADMIN_URL envar")

	KratosPublicURL := MustMakeURLValue(os.Getenv("KRATOS_PUBLIC_URL"))
	flag.Var(&KratosPublicURL, "kratos-public-url", "The URL where ORY Kratos's Public API is located at. If this app and ORY Kratos are running in the same private network, this should be the private network address. Defaults to KRATOS_PUBLIC_URL envar")

	KratosBrowserURL := MustMakeURLValue(os.Getenv("KRATOS_BROWSER_URL"))
	flag.Var(&KratosBrowserURL, "kratos-browser-url", "The URL to build all of the kratos self service URLS. Defaults to KRATOS_BROWSER_URL envar")

	BaseURL := MustMakeURLValue(os.Getenv("BASE_URL"))
	flag.Var(&BaseURL, "base-url", "The base url of this app. If served e.g. behind a proxy or via GitHub pages this would be the path, e.g. https://mywebsite.com/kratos-selfservice-ui-go/. Must be absolute!. Defaults to BASE_URL envar")

	flag.StringVar(&o.Host, "host", "0.0.0.0", "Optional host that app listens on.")

	flag.IntVar(&o.Port, "port", parseInt(os.Getenv("PORT")), "Port for this app to listen on. Defaults to PORT envar")

	flag.DurationVar(&o.ShutdownWait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")

	flag.StringVar(&o.TLSCertPath, "tls-cert-path", "", "Optional path to the certificate file. Use in conjunction with tls-key-path to enable https.")

	flag.StringVar(&o.TLSKeyPath, "tls-key-path", "", "Optional path to the key file. Use in conjunction with tls-cert-path to enable https.")

	var allCookieStoreKeyPairs string
	flag.StringVar(&allCookieStoreKeyPairs, "cookie-store-key-pairs", os.Getenv("COOKIE_STORE_KEY_PAIRS"), "Pairs of authentication and encryption keys, enclose then in quotes. See the gen-cookie-store-key-pair flag to generate")

	genCookieStoreKeys := false
	flag.BoolVar(&genCookieStoreKeys, "gen-cookie-store-key-pair", false, "Pass this flag to generate a pairs of authentication and encryption keys and exit")

	flag.Parse()

	if genCookieStoreKeys {
		authKey := securecookie.GenerateRandomKey(32)
		encrKey := securecookie.GenerateRandomKey(32)
		fmt.Printf("The following values are suitable for passing in 'cookie-store-key-pairs' flag:\n")
		fmt.Printf("%s %s\n",
			base64.StdEncoding.EncodeToString(authKey),
			base64.StdEncoding.EncodeToString(encrKey))
		os.Exit(0)
	}

	o.KratosAdminURL = KratosAdminURL.URL
	o.KratosPublicURL = KratosPublicURL.URL
	o.KratosBrowserURL = KratosBrowserURL.URL
	o.BaseURL = BaseURL.URL
	o.CookieStoreKeyPairs = make([][]byte, 0)
	pairs := strings.Split(allCookieStoreKeyPairs, " ")
	for _, s := range pairs {
		decoded, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			log.Fatalf("Error decoding 'cookie-store-key-pairs' value: '%s' , did you use 'gen-cookie-store-key-pair' to generate them? Error: %v", s, err)
		}
		o.CookieStoreKeyPairs = append(o.CookieStoreKeyPairs, []byte(decoded))
	}
	return o
}

// Validate checks that the options are valid and return nil, or returns an error
func (o *Options) Validate() error {

	if o.KratosAdminURL == nil || o.KratosAdminURL.String() == "" {
		return errors.New("'kratos-admin-url' URL missing")
	}

	if o.KratosPublicURL == nil || o.KratosPublicURL.String() == "" {
		return errors.New("'kratos-public-url' URL missing")
	}

	if o.KratosBrowserURL == nil || o.KratosBrowserURL.String() == "" {
		return errors.New("'kratos-browser-url' URL missing")
	}

	if o.BaseURL == nil || o.BaseURL.String() == "" {
		return errors.New("'base-url' URL missing")
	}

	if o.TLSCertPath != "" && !fileExists(o.TLSCertPath) {
		return fmt.Errorf("'tls-cert-path' file '%s' invalid", o.TLSCertPath)
	}

	if o.TLSKeyPath != "" && !fileExists(o.TLSKeyPath) {
		return fmt.Errorf("'tls-key-path' file '%s' invalid", o.TLSKeyPath)
	}

	if (o.TLSCertPath == "" && o.TLSKeyPath != "") || (o.TLSCertPath != "" && o.TLSKeyPath == "") {
		return fmt.Errorf("To enable HTTPS, provide 'tls-key-path' and 'tls-cert-path'")
	}

	if !(len(o.CookieStoreKeyPairs) == 1 || len(o.CookieStoreKeyPairs)%2 == 0) {
		return fmt.Errorf("'cookie-store-key-pairs' has %d values, it should contain one auth key, or even pairs of auth & encryption keys separated by a space", len(o.CookieStoreKeyPairs))
	}

	return nil
}

// WhoAmIURL returns the URL to POST to to get the session
func (o *Options) WhoAmIURL() string {
	url := o.KratosPublicURL
	url.Path = "/sessions/whoami"
	return url.String()
}

// RegistrationURL returns the URL to redirect to that will
// start the registration flow
func (o *Options) RegistrationURL() string {
	url := o.KratosBrowserURL
	url.Path = "/self-service/registration/browser"
	return url.String()
}

// SettingsURL returns the URL to redirect to that will
// start the settings flow
func (o *Options) SettingsURL() string {
	url := o.KratosBrowserURL
	url.Path = "/self-service/settings/browser"
	return url.String()
}

// LoginFlowURL returns the URL to redirect to that will
// start the login flow
func (o *Options) LoginFlowURL() string {
	url := o.KratosBrowserURL
	url.Path = "/self-service/login/browser"
	return url.String()
}

// RecoveryFlowURL returns the URL to redirect to that will
// start the recovery flow
func (o *Options) RecoveryFlowURL() string {
	url := o.KratosBrowserURL
	url.Path = "/self-service/recovery/browser"
	return url.String()
}

// LogoutFlowURL returns the URL to redirect to that will
// start the logout flow
func (o *Options) LogoutFlowURL() string {
	url := o.KratosBrowserURL
	url.Path = "/self-service/browser/flows/logout"
	return url.String()
}

// LoginURL returns the URL to redirect to that shows the login page
func (o *Options) LoginPageURL() string {
	url := o.BaseURL
	url.Path = "/auth/login"
	return url.String()
}

// Address that this application will listen on
func (o *Options) Address() string {
	return fmt.Sprintf("%s:%d", o.Host, o.Port)
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
