package options

import (
	"io/ioutil"
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
	assert.Nil(t, o.Validate())

	// If provide key or cert, must have both
	o = Options{
		TLSCertPath: certFile.Name(),
		TLSKeyPath:  "",
	}
	assert.EqualError(t, o.Validate(), "To enable HTTPS, provide 'tls-key-path' and 'tls-cert-path'")

}
