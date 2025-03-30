package main

import (
	"net/http"
	"net/http/httptest"
	"project/orm/model"
	"project/route"
	"strings"
	"testing"

	"encoding/json"

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

func TestPostUser(t *testing.T) {
	router := route.SetupRouter()
	router = route.PostUser(router)

	w := httptest.NewRecorder()

	exampleUser := model.User{
		Username: "exampleUser",
	}
	userJson, _ := json.Marshal(exampleUser)
	req, _ := http.NewRequest("POST", "/user/add", strings.NewReader(string(userJson)))
	router.ServeHTTP(w, req)

	res := string(`{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"Username":"exampleUser"}`)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, res, w.Body.String())
}
