package xerrors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFoundErrorError(t *testing.T) {
	err := NotFoundError{
		Code:        "404",
		Description: " resource not found",
	}
	expected := "Not Found - 404 resource not found"
	assert.Equal(t, expected, err.Error())
}

func TestNewNotFoundError(t *testing.T) {
	code := "404"
	desc := " missing item"
	err := NewNotFoundError(code, desc)
	assert.Equal(t, code, err.Code)
	assert.Equal(t, desc, err.Description)
	assert.Equal(t, "Not Found - 404 missing item", err.Error())
}
