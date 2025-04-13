package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hedagaurav/blog-api/config"
	"github.com/hedagaurav/blog-api/models"
	"net/http"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.ConnectDatabase().Create(&post)
	c.JSON(http.StatusOK, post)
}

// GetPosts retrieves all posts
func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := config.ConnectDatabase().Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// GetPost retrieves a single post by ID
func GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := config.ConnectDatabase().First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// UpdatePost updates a post by ID
func UpdatePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := config.ConnectDatabase().First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.ConnectDatabase().Save(&post)
	c.JSON(http.StatusOK, post)
}

// DeletePost deletes a post by ID
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	if err := config.ConnectDatabase().Delete(&models.Post{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
