package ulid

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUlid(t *testing.T) {
	id := NewUlid()
	assert.NotEqual(t, uuid.Nil, id)
	_, err := uuid.Parse(id.String())
	assert.NoError(t, err)
}

func TestULIDFromStringValid(t *testing.T) {
	id := NewUlid()
	parsed, err := ULIDFromString(id.String())
	assert.NoError(t, err)
	assert.Equal(t, id, parsed)
}

func TestULIDFromStringInvalid(t *testing.T) {
	_, err := ULIDFromString("invalid-ulid-string")
	assert.Error(t, err)
}
