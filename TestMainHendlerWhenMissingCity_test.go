package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenMissingCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.NotEqual(t, responseRecorder.Code, http.StatusBadRequest)

	expected := `wrong city value`
	assert.NotEqual(t, responseRecorder.Body.String(), expected)
}
