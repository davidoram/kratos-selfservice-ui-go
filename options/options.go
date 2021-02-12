package options

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

type Options struct {
	// KratosAdminURL is the URL where ORY Kratos's Admin API is located at. If this app and ORY Kratos are running in the same private network, this should be the private network address (e.g. kratos-admin.svc.cluster.local).
	KratosAdminURL url.URL

	// KratosPublicURL is the URL where ORY Kratos's Public API is located at. If this app and ORY Kratos are running in the same private network, this should be the private network address (e.g. kratos-public.svc.cluster.local).
	KratosPublicURL url.URL

	// BaseURL is the base url of this app. If served e.g. behind a proxy or via GitHub pages this would be the path, e.g. https://mywebsite.com/kratos-selfservice-ui-go/. Must be absolute!
	BaseURL url.URL

	// Port that this app is listening on
	Port int

	// TLSCertPath is an optional Path to certificate file. Should be set up together with TLSKeyPath to enable HTTPS.
	TLSCertPath string
	// TLSCertPath is an optional path to key file Should be set up together with TLSCertPath to enable HTTPS.
	TLSKeyPath string
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) SetFromCommandLine() *Options {

	flag.Var(&URLValue{&o.KratosAdminURL, ""}, "kratos-admin-url", "The URL where ORY Kratos's Admin API is located at. If this app and ORY Kratos are running in the same private network, this should be the private network address.")

	flag.Var(&URLValue{&o.KratosPublicURL, ""}, "kratos-public-url", "The URL where ORY Kratos's Public API is located at. If this app and ORY Kratos are running in the same private network, this should be the private network address.")

	defaultBaseURL := "/"
	flag.Var(&URLValue{&o.BaseURL, defaultBaseURL},
		"base-url",
		fmt.Sprintf("The base url of this app. If served e.g. behind a proxy or via GitHub pages this would be the path, e.g. https://mywebsite.com/kratos-selfservice-ui-go/. Must be absolute!. Defaults to '%s'", defaultBaseURL))

	flag.IntVar(&o.Port, "port", 4455, "Port for this app to listen on.")

	flag.StringVar(&o.TLSCertPath, "tls-cert-path", "", "Optional path to the certificate file. Use in conjunction with tls-key-path to enable https.")

	flag.StringVar(&o.TLSKeyPath, "tls-key-path", "", "Optional path to the key file. Use in conjunction with tls-cert-path to enable https.")
	flag.Parse()

	return o
}

// Validate checks that the options are valid and return nil, or returns an error
func (o *Options) Validate() error {

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

// Address that this application will listen on
func (o *Options) Address() string {
	return fmt.Sprintf(":%d", o.Port)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
