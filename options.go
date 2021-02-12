package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

type Options struct {
	// KratosAdminURL is the URL where ORY Kratos's Admin API is located at. If this app and ORY Kratos are running in the same private network, this should be the private network address (e.g. kratos-admin.svc.cluster.local).
	KratosAdminURL string

	// KratosPublicURL is the URL where ORY Kratos's Public API is located at. If this app and ORY Kratos are running in the same private network, this should be the private network address (e.g. kratos-public.svc.cluster.local).
	KratosPublicURL string

	// BaseURL is the base url of this app. If served e.g. behind a proxy or via GitHub pages this would be the path, e.g. https://mywebsite.com/kratos-selfservice-ui-go/. Must be absolute!
	BaseURL string

	// TLSCertPath is an optional Path to certificate file. Should be set up together with TLSKeyPath to enable HTTPS.
	TLSCertPath string
	// TLSCertPath is an optional path to key file Should be set up together with TLSCertPath to enable HTTPS.
	TLSKeyPath string
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) SetFromCommandLine() *Options {

	flag.StringVar(&o.KratosAdminURL, "kratos-admin-url", "", "The URL where ORY Kratos's Admin API is located at. If this app and ORY Kratos are running in the same private network, this should be the private network address.")

	flag.StringVar(&o.KratosPublicURL, "kratos-public-url", "", "The URL where ORY Kratos's Public API is located at. If this app and ORY Kratos are running in the same private network, this should be the private network address.")

	flag.StringVar(&o.BaseURL, "base-url", "", "The base url of this app. If served e.g. behind a proxy or via GitHub pages this would be the path, e.g. https://mywebsite.com/kratos-selfservice-ui-go/. Must be absolute!")

	flag.StringVar(&o.TLSCertPath, "tls-cert-path", "", "Optional path to the certificate file. Use in conjunction with tls-key-path to enable https.")

	flag.StringVar(&o.TLSKeyPath, "tls-key-path", "", "Optional path to the key file. Use in conjunction with tls-cert-path to enable https.")
	flag.Parse()

	return o
}

// Validate checks that the options are valid and return nil, or returns an error
func (o *Options) Validate() error {

	if _, err := url.Parse(o.KratosAdminURL); err != nil {
		return fmt.Errorf("'kratos-admin-url' URL '%s' invalid: %v", o.KratosAdminURL, err)
	}

	if _, err := url.Parse(o.KratosPublicURL); err != nil {
		return fmt.Errorf("'kratos-public-url' URL '%s' invalid: %v", o.KratosPublicURL, err)
	}

	if _, err := url.Parse(o.BaseURL); err != nil {
		return fmt.Errorf("'base-url' URL '%s' invalid: %v", o.BaseURL, err)
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

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
