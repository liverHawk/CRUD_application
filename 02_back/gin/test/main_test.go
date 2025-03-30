package test

import (
	"net/http"
	"net/http/httptest"
	"project/route"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := route.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	res := string(`{"message":"pong"}`)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, res, w.Body.String())
}
