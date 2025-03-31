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

func (h *Handler) GetUser(r *gin.Engine) *gin.Engine {
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user model.User
		if err := h.DB.First(&user, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(200, user)
	})
	return r
}

func (h *Handler) UpdateUser(r *gin.Engine) *gin.Engine {
	r.PUT("/user/update/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
		if err := h.DB.Model(&model.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update user"})
			return
		}
		c.JSON(200, user)
	})
	return r
}

func (h *Handler) DeleteUser(r *gin.Engine) *gin.Engine {
	r.DELETE("/user/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := h.DB.Delete(&model.User{}, id).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete user"})
			return
		}
		c.JSON(200, gin.H{"message": "User deleted successfully"})
	})
	return r
}

func (h *Handler) UserRoutes(r *gin.Engine) *gin.Engine {
	r = h.PostUser(r)
	r = h.GetUser(r)
	r = h.UpdateUser(r)
	r = h.DeleteUser(r)

	return r
}
