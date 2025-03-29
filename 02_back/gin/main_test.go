package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	// "encoding/json"
	// "strings"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	res := string(`{"message":"pong"}`)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, res, w.Body.String())
}
