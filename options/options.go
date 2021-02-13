package options

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

type Options struct {
	// KratosAdminURL is the URL where ORY Kratos's Admin API is located at. If this app and ORY Kratos are running in the same private network, this should be the private network address (e.g. kratos-admin.svc.cluster.local).
	KratosAdminURL *url.URL

	// KratosPublicURL is the URL where ORY Kratos's Public API is located at. If this app and ORY Kratos are running in the same private network, this should be the private network address (e.g. kratos-public.svc.cluster.local).
	KratosPublicURL *url.URL

	// KratosBrowserURL is the URL where ORY Kratos's self service browser endpoints are located at.
	KratosBrowserURL *url.URL

	// BaseURL is the base url of this app. If served e.g. behind a proxy or via GitHub pages this would be the path, e.g. https://mywebsite.com/kratos-selfservice-ui-go/. Must be absolute!
	BaseURL *url.URL

	// Port that this app is listening on
	Port int

	// TLSCertPath is an optional Path to certificate file. Should be set up together with TLSKeyPath to enable HTTPS.
	TLSCertPath string
	// TLSCertPath is an optional path to key file Should be set up together with TLSCertPath to enable HTTPS.
	TLSKeyPath string
}

func NewOptions() *Options {
	return &Options{
		KratosAdminURL:   &url.URL{},
		KratosPublicURL:  &url.URL{},
		KratosBrowserURL: &url.URL{},
		BaseURL:          &url.URL{},
	}
}

func (o *Options) SetFromCommandLine() *Options {

	KratosAdminURL := MustMakeURLValue(os.Getenv("KRATOS_ADMIN_URL"))
	flag.Var(&KratosAdminURL, "kratos-admin-url", "The URL where ORY Kratos's Admin API is located at. If this app and ORY Kratos are running in the same private network, this should be the private network address. Defaults to KRATOS_ADMIN_URL envar")

	KratosPublicURL := MustMakeURLValue(os.Getenv("KRATOS_PUBLIC_URL"))
	flag.Var(&KratosPublicURL, "kratos-public-url", "The URL where ORY Kratos's Public API is located at. If this app and ORY Kratos are running in the same private network, this should be the private network address. Defaults to KRATOS_PUBLIC_URL envar")

	KratosBrowserURL := MustMakeURLValue(os.Getenv("KRATOS_BROWSER_URL"))
	flag.Var(&KratosBrowserURL, "kratos-browser-url", "The URL to build all of the kratos self service URLS. Defaults to KRATOS_BROWSER_URL envar")

	BaseURL := MustMakeURLValue(os.Getenv("BASE_URL"))
	flag.Var(&BaseURL, "base-url", "The base url of this app. If served e.g. behind a proxy or via GitHub pages this would be the path, e.g. https://mywebsite.com/kratos-selfservice-ui-go/. Must be absolute!. Defaults to BASE_URL envar")

	flag.IntVar(&o.Port, "port", parseInt(os.Getenv("PORT")), "Port for this app to listen on. Defaults to PORT envar")

	flag.StringVar(&o.TLSCertPath, "tls-cert-path", "", "Optional path to the certificate file. Use in conjunction with tls-key-path to enable https.")

	flag.StringVar(&o.TLSKeyPath, "tls-key-path", "", "Optional path to the key file. Use in conjunction with tls-cert-path to enable https.")
	flag.Parse()

	o.KratosAdminURL = KratosAdminURL.URL
	o.KratosPublicURL = KratosPublicURL.URL
	o.KratosBrowserURL = KratosBrowserURL.URL
	o.BaseURL = BaseURL.URL

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

	return nil
}

// RegistrationURL returns the URL to redirect to that will
// start the registration flow
func (o *Options) RegistrationURL() string {
	url := o.KratosPublicURL
	url.Path = "/self-service/registration/browser"
	return url.String()
}

// LoginFlowURL returns the URL to redirect to that will
// start the login flow
func (o *Options) LoginFlowURL() string {
	url := o.KratosPublicURL
	url.Path = "/self-service/login/browser"
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
	return fmt.Sprintf(":%d", o.Port)
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
