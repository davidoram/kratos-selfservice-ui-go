package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateOptions(t *testing.T) {
	o := Options{
		KratosAdminURL:  "http://kratos-admin",
		KratosPublicURL: "http://kratos-public",
		BaseURL:         "https://mysite.com",
	}
	assert.Nil(t, o.Validate())

	// Error on invalid URL
	o = Options{
		KratosAdminURL: ":http//kratos-admin",
	}
	assert.EqualError(t, o.Validate(), "'kratos-admin-url' URL ':http//kratos-admin' invalid: parse \":http//kratos-admin\": missing protocol scheme")

	// Temporary files
	certFile, err := ioutil.TempFile("", "certFile")
	assert.Nil(t, err)
	defer os.Remove(certFile.Name())

	keyFile, err := ioutil.TempFile("", "keyFile")
	assert.Nil(t, err)
	defer os.Remove(keyFile.Name())

	o = Options{
		TLSCertPath: certFile.Name(),
		TLSKeyPath:  keyFile.Name(),
	}
	assert.Nil(t, o.Validate())

	// If provide key or cert, must have both
	o = Options{
		TLSCertPath: certFile.Name(),
		TLSKeyPath:  "",
	}
	assert.EqualError(t, o.Validate(), "To enable HTTPS, provide 'tls-key-path' and 'tls-cert-path'")

}
