package options

import (
	"io/ioutil"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateOptions(t *testing.T) {
	// Temporary files
	certFile, err := ioutil.TempFile("", "certFile")
	assert.Nil(t, err)
	defer os.Remove(certFile.Name())

	keyFile, err := ioutil.TempFile("", "keyFile")
	assert.Nil(t, err)
	defer os.Remove(keyFile.Name())

	o := Options{
		TLSCertPath: certFile.Name(),
		TLSKeyPath:  keyFile.Name(),
	}

	// Check URLs must be supplied
	assert.EqualError(t, o.Validate(), "'kratos-admin-url' URL missing")

	u, _ := url.Parse("http://testhost.com")
	o.KratosAdminURL = u
	assert.EqualError(t, o.Validate(), "'kratos-public-url' URL missing")
	o.KratosPublicURL = u
	assert.EqualError(t, o.Validate(), "'kratos-browser-url' URL missing")
	o.KratosBrowserURL = u
	assert.EqualError(t, o.Validate(), "'base-url' URL missing")
	o.BaseURL = u

	// If provide key or cert, must have both
	o.TLSKeyPath = ""
	assert.EqualError(t, o.Validate(), "To enable HTTPS, provide 'tls-key-path' and 'tls-cert-path'")

	// File paths must be valid
	o.TLSKeyPath = "/not/a/valid/path"
	assert.EqualError(t, o.Validate(), "'tls-key-path' file '/not/a/valid/path' invalid")

}
