package route

import (
	"project/orm/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateArticle(r *gin.Engine) *gin.Engine {
	r.POST("/article/add", func(c *gin.Context) {
		id := c.Query("id")
		if id == "" {
			c.JSON(400, gin.H{"error": "User ID is required"})
			return
		}

		// Check if the user exists before creating an article
		if !model.ExistUser(h.DB, id) {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		var article model.Article
		if err := c.ShouldBindJSON(&article); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		// id:string -> uint -> article.AuthorID
		authorID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		article.AuthorID = uint(authorID)

		statusCode := model.CreateArticle(h.DB, &article)
		if statusCode == 500 {
			c.JSON(500, gin.H{"error": "Failed to create article"})
			return
		}
		c.JSON(200, article)
	})
	return r
}

func (h *Handler) GetArticle(r *gin.Engine) *gin.Engine {
	r.GET("/article/:id", func(c *gin.Context) {
		id := c.Param("id")

		if !model.ExistArticle(h.DB, id) {
			c.JSON(404, gin.H{"error": "Article not found"})
			return
		}

		var article model.Article
		statusCode := model.GetArticle(h.DB, &article, id)
		if statusCode == 500 {
			c.JSON(500, gin.H{"error": "Failed to get article"})
			return
		}
		c.JSON(200, article)
	})
	return r
}

func (h *Handler) UpdateArticle(r *gin.Engine) *gin.Engine {
	r.PUT("/article/update/:id", func(c *gin.Context) {
		authorID := c.Query("id")
		if authorID == "" {
			c.JSON(400, gin.H{"error": "User ID is required"})
			return
		}

		if !model.ExistUser(h.DB, authorID) {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		id := c.Param("id")

		if !model.ExistArticle(h.DB, id) {
			c.JSON(404, gin.H{"error": "Article not found"})
			return
		}

		var article model.Article
		if err := c.ShouldBindJSON(&article); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		statusCode := model.UpdateArticle(h.DB, &article, id)
		if statusCode == 500 {
			c.JSON(500, gin.H{"error": "Failed to update article"})
			return
		}
		c.JSON(200, article)
	})
	return r
}

func (h *Handler) DeleteArticle(r *gin.Engine) *gin.Engine {
	r.DELETE("/article/delete/:id", func(c *gin.Context) {
		authorID := c.Query("id")
		if authorID == "" {
			c.JSON(400, gin.H{"error": "User ID is required"})
			return
		}

		if !model.ExistUser(h.DB, authorID) {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		id := c.Param("id")

		if !model.ExistArticle(h.DB, id) {
			c.JSON(404, gin.H{"error": "Article not found"})
			return
		}

		statusCode := model.DeleteArticle(h.DB, id)
		if statusCode == 500 {
			c.JSON(500, gin.H{"error": "Failed to delete article"})
			return
		}
		c.JSON(200, gin.H{"message": "Article deleted successfully"})
	})
	return r
}

func (h *Handler) ArticleRoutes(r *gin.Engine) *gin.Engine {
	r = h.CreateArticle(r)
	r = h.GetArticle(r)
	r = h.UpdateArticle(r)
	r = h.DeleteArticle(r)

	return r
}
