package validator

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/tbtec/tremligeiro/internal/types/xerrors"
)

type testStruct struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
    Age   int    `validate:"min=18"`
}

func TestValidate_Valid(t *testing.T) {
    input := testStruct{Name: "Alice", Email: "alice@email.com", Age: 20}
    err := Validate(input)
    assert.NoError(t, err)
}

func TestValidate_MissingRequired(t *testing.T) {
    input := testStruct{Email: "alice@email.com", Age: 20}
    err := Validate(input)
    assert.Error(t, err)
    vErr, ok := err.(xerrors.ValidationError)
    assert.True(t, ok)
    assert.NotEmpty(t, vErr.Fields)
    found := false
    for _, f := range vErr.Fields {
        if f.Name == "Name" {
            found = true
            assert.Contains(t, f.Reasons, xerrors.ReasonRequiredAttributeMissing)
        }
    }
    assert.True(t, found)
}

func TestValidate_InvalidEmailAndAge(t *testing.T) {
    input := testStruct{Name: "Alice", Email: "not-an-email", Age: 10}
    err := Validate(input)
    assert.Error(t, err)
    vErr, ok := err.(xerrors.ValidationError)
    assert.True(t, ok)
    assert.NotEmpty(t, vErr.Fields)
    var emailFound, ageFound bool
    for _, f := range vErr.Fields {
        if f.Name == "Email" {
            emailFound = true
            assert.Contains(t, f.Reasons, xerrors.ReasonTypeInvalidValue)
        }
        if f.Name == "Age" {
            ageFound = true
            assert.Contains(t, f.Reasons, xerrors.ReasonTypeInvalidValue)
        }
    }
    assert.True(t, emailFound)
    assert.True(t, ageFound)
}

func TestExtract(t *testing.T) {
    assert.Equal(t, "Field", extract("Struct.Field"))
    assert.Equal(t, "Field.Sub", extract("Struct.Field.Sub"))
    assert.Equal(t, "Field", extract("Field"))
}