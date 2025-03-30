package route

import (
	"project/orm/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostUser(r *gin.Engine) *gin.Engine {
	r.POST("/user/add", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
		if err := h.DB.Create(&user).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to create user"})
			return
		}
		c.JSON(200, user)
	})
	return r
}
