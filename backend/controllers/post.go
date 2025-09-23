package controllers

import (
	"flower-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	ErrFlowerNotFound = "Flower not found"
	ErrUserNotFound   = "User not found"
)

// Get /api/flowers
func ListFlowers(c *gin.Context) {
	postService := services.NewPostService()

	flowers, err := postService.GetAllPostsWithAuthor()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query flowers"})
		return
	}

	c.JSON(http.StatusOK, flowers)
}

// Get /api/flowers/:id
func GetFlower(c *gin.Context) {
	id := c.Param("id")
	postService := services.NewPostService()

	flower, err := postService.GetPostWithAuthorByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrFlowerNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query flower"})
		return
	}

	c.JSON(http.StatusOK, flower)
}

// Post /api/flowers
func CreateFlower(c *gin.Context) {
	var payload struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		ImageURL string `json:"image_url"`
	}
	if err := c.Bind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	uid, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var userId uint
	switch v := uid.(type) {
	case float64:
		userId = uint(v)
	case int:
		userId = uint(v)
	case int64:
		userId = uint(v)
	case uint:
		userId = v
	}

	postService := services.NewPostService()
	flower, err := postService.CreatePost(payload.Title, payload.Content, payload.ImageURL, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create flower"})
		return
	}

	c.JSON(http.StatusCreated, flower)
}

// Put /api/flowers/:id
func UpdateFlower(c *gin.Context) {
	idStr := c.Param("id")

	var payload struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		ImageURL string `json:"image_url"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	uid, _ := c.Get("userId")
	role, _ := c.Get("role")

	var userId uint
	switch v := uid.(type) {
	case float64:
		userId = uint(v)
	case int:
		userId = uint(v)
	case int64:
		userId = uint(v)
	case uint:
		userId = v
	}

	postService := services.NewPostService()

	// Check ownership
	isOwner, err := postService.CheckPostOwnership(idStr, userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrFlowerNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check ownership"})
		return
	}

	isAdmin := false
	if rs, ok := role.(string); ok && rs == "admin" {
		isAdmin = true
	}

	if !isOwner && !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	flower, err := postService.UpdatePost(idStr, payload.Title, payload.Content, payload.ImageURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update flower"})
		return
	}

	c.JSON(http.StatusOK, flower)
}

// Delete /api/flowers/:id
func DeleteFlower(c *gin.Context) {
	idStr := c.Param("id")

	uid, _ := c.Get("userId")
	role, _ := c.Get("role")

	var userId uint
	switch v := uid.(type) {
	case float64:
		userId = uint(v)
	case int:
		userId = uint(v)
	case int64:
		userId = uint(v)
	case uint:
		userId = v
	}

	postService := services.NewPostService()

	// Check ownership
	isOwner, err := postService.CheckPostOwnership(idStr, userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrFlowerNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check ownership"})
		return
	}

	isAdmin := false
	if rs, ok := role.(string); ok && rs == "admin" {
		isAdmin = true
	}

	if !isOwner && !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	if err := postService.DeletePost(idStr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete flower"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flower deleted successfully"})
}

func LikeFlower(c *gin.Context) {
	idStr := c.Param("id")
	postService := services.NewPostService()

	flower, err := postService.LikePost(idStr)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrFlowerNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like flower"})
		return
	}

	c.JSON(http.StatusOK, flower)
}

func UnlikeFlower(c *gin.Context) {
	idStr := c.Param("id")
	postService := services.NewPostService()

	flower, err := postService.UnlikePost(idStr)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrFlowerNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike flower"})
		return
	}

	c.JSON(http.StatusOK, flower)
}

// Get /api/user
func GetUser(c *gin.Context) {
	uid, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var userId uint
	switch v := uid.(type) {
	case float64:
		userId = uint(v)
	case int:
		userId = uint(v)
	case int64:
		userId = uint(v)
	case uint:
		userId = v
	}

	userService := services.NewUserService()
	user, err := userService.GetUserByID(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrUserNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
