package httpclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := New()
	assert.NotNil(t, client)
}
