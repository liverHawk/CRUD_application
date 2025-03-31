package test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"project/orm/model"
	"project/route"
	"project/util"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestPostUser(t *testing.T) {
	driverName := fmt.Sprintf("test_post_user_%d", time.Now().UnixNano())

	db, _ := util.NewTestDB(driverName)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

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

func postUser(u *model.User, r *gin.Engine) *model.User {
	body, _ := json.Marshal(u)
	req, _ := http.NewRequest("POST", "/user/add", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseUser model.User
	err := json.Unmarshal(w.Body.Bytes(), &responseUser)
	if err != nil {
		return nil
	}
	return &responseUser
}

func createUser(db *gorm.DB, u *model.User) *model.User {
	if err := db.Create(u).Error; err != nil {
		return nil
	}
	return u
}

func TestGetUser(t *testing.T) {
	driverName := fmt.Sprintf("test_post_user_%d", time.Now().UnixNano())

	db, _ := util.NewTestDB(driverName)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()
	h := route.Handler{
		DB: db,
	}

	router := route.SetupRouter()
	router = h.GetUser(router)
	router = h.PostUser(router)

	w := httptest.NewRecorder()

	exampleUser := model.User{
		Username: "exampleUser",
	}
	var createdUser model.User
	createdUser = *postUser(&exampleUser, router)

	req, _ := http.NewRequest("GET", "/user/"+strconv.FormatUint(uint64(createdUser.ID), 10), nil)
	router.ServeHTTP(w, req)

	var responseUser model.User
	err := json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.NoError(t, err, "Failed to unmarshal response")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, exampleUser.Username, responseUser.Username, "Username should match")
}

func TestUpdateUser(t *testing.T) {
	driverName := fmt.Sprintf("test_post_user_%d", time.Now().UnixNano())

	db, _ := util.NewTestDB(driverName)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()
	h := route.Handler{
		DB: db,
	}

	router := route.SetupRouter()
	// otherRouter := h.PostUser(router)
	router = h.UpdateUser(router)

	w := httptest.NewRecorder()

	exampleUser := model.User{
		Username: "exampleUser",
	}
	var createdUser model.User
	// createdUser = *postUser(&exampleUser, otherRouter)
	createdUser = *createUser(db, &exampleUser)

	time.Sleep(50 * time.Millisecond)

	reqBody := model.User{
		Username: "updatedUser",
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("PUT", "/user/update/"+strconv.FormatUint(uint64(createdUser.ID), 10), strings.NewReader(string(reqBodyJson)))
	router.ServeHTTP(w, req)

	var responseUser model.User
	log.Println(w.Body.String())
	err := json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.NoError(t, err, "Failed to unmarshal response")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "updatedUser", responseUser.Username, "Username should be updated")
}

func TestDeleteUser(t *testing.T) {
	driverName := fmt.Sprintf("test_post_user_%d", time.Now().UnixNano())

	db, _ := util.NewTestDB(driverName)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()
	h := route.Handler{
		DB: db,
	}

	router := route.SetupRouter()
	router = h.DeleteUser(router)

	router = h.PostUser(router)

	w := httptest.NewRecorder()

	exampleUser := model.User{
		Username: "exampleUser",
	}
	var createdUser model.User
	createdUser = *postUser(&exampleUser, router)

	req, _ := http.NewRequest("DELETE", "/user/delete/"+strconv.FormatUint(uint64(createdUser.ID), 10), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Check if the user is actually deleted
	var deletedUser model.User
	err := db.First(&deletedUser, exampleUser.ID).Error
	if err == nil {
		t.Errorf("User with ID %d should have been deleted", exampleUser.ID)
	}
}
