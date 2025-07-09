package httpserver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestBuilderBuild(t *testing.T) {
	headers := map[string]string{"Authorization": "Bearer token"}
	params := map[string]string{"id": "123"}
	query := map[string]string{"q": "search"}
	body := []byte(`{"foo":"bar"}`)

	req := NewRequestBuilder().
		Host("localhost").
		Path("/api/v1/resource").
		Method("POST").
		Headers(headers).
		Params(params).
		Query(query).
		Body(body).
		Build()

	assert.Equal(t, "localhost", req.Host)
	assert.Equal(t, "/api/v1/resource", req.Path)
	assert.Equal(t, "POST", req.Method)
	assert.Equal(t, headers, req.Headers)
	assert.Equal(t, params, req.Params)
	assert.Equal(t, query, req.Query)
	assert.Equal(t, body, req.Body)
}
