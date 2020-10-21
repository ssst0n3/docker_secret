package cert

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateCertificate(t *testing.T) {
	ca, key, err := GenerateCertificate()
	assert.NoError(t, err)
	assert.Equal(t, true, bytes.Contains(ca, []byte("CERTIFICATE")))
	assert.Equal(t, true, bytes.Contains(key, []byte("RSA PRIVATE KEY")))
}
