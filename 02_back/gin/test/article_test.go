package test

import (
	"fmt"
	"project/orm/model"
	"project/route"
	"project/util"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func beforeAction(t *testing.T) (*gin.Engine, route.Handler) {
	driverName := fmt.Sprintf("test_%s_%d", "article", time.Now().UnixNano())
	db, _ := util.NewTestDB(driverName)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	h := route.Handler{
		DB: db,
	}
	exampleUser := model.User{
		Username: "exampleUser",
	}

	u := createUser(db, &exampleUser)
	if u == nil {
		assert.NotNil(t, u, "Failed to create user")
	}

	router := route.SetupRouter()
	return router, h
}

func TestCreateArticle(t *testing.T) {
	router, h := beforeAction(t)

	router = h.CreateArticle(router)

}
