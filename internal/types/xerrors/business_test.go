package xerrors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBusinessErrorError(t *testing.T) {
	err := BusinessError{
		Code:        "123",
		Description: "some error",
	}
	expected := "123 - some error"
	assert.Equal(t, expected, err.Error())
}

func TestNewBusinessError(t *testing.T) {
	code := "456"
	desc := "another error"
	err := NewBusinessError(code, desc)
	assert.Equal(t, code, err.Code)
	assert.Equal(t, desc, err.Description)
	assert.Equal(t, code+" - "+desc, err.Error())
}
