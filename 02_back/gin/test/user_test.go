package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project/orm/model"
	"project/route"
	"project/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostUser(t *testing.T) {
	db, _ := util.NewTestDB("test")
	h := route.Handler{
		DB: db,
	}

	router := route.SetupRouter()
	router = h.PostUser(router)

	w := httptest.NewRecorder()

	exampleUser := model.User{
		Username: "exampleUser",
	}
	userJson, _ := json.Marshal(exampleUser)
	req, _ := http.NewRequest("POST", "/user/add", strings.NewReader(string(userJson)))
	router.ServeHTTP(w, req)

	var responseUser model.User
	err := json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.NoError(t, err, "Failed to unmarshal response")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, exampleUser.Username, responseUser.Username, "Username should match")
	assert.NotZero(t, responseUser.ID, "ID should not be zero")
}
