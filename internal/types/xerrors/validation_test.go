package xerrors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValidationError(t *testing.T) {
	err := NewValidationError("invalid input")
	assert.Equal(t, "invalid input", err.Description)
	assert.Empty(t, err.Fields)
}

func TestValidationErrorAddField(t *testing.T) {
	err := NewValidationError("invalid input")
	err = err.AddField("field1", ReasonTypeInvalidValue)
	err = err.AddField("field2", ReasonRequiredAttributeMissing, "CUSTOM_REASON")

	assert.Len(t, err.Fields, 2)
	assert.Equal(t, "field1", err.Fields[0].Name)
	assert.Equal(t, []string{ReasonTypeInvalidValue}, err.Fields[0].Reasons)
	assert.Equal(t, "field2", err.Fields[1].Name)
	assert.Equal(t, []string{ReasonRequiredAttributeMissing, "CUSTOM_REASON"}, err.Fields[1].Reasons)
}

func TestValidationErrorError(t *testing.T) {
	err := NewValidationError("desc")
	assert.Equal(t, "0 fields are invalid", err.Error())

	err = err.AddField("field1", ReasonTypeInvalidValue)
	assert.Equal(t, "1 fields are invalid", err.Error())
}
