package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewNotFound(t *testing.T) {
	app := fiber.New()
	app.Use(NewNotFound())

	req := httptest.NewRequest("GET", "/non-existent-path", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
