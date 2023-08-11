package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHello(t *testing.T) {
	api := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	api.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}