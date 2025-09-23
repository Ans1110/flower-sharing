package controllers

import (
	"flower-backend/database"
	"flower-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get /api/flowers
func ListFlowers(c *gin.Context) {
	var flowers []models.Post
	if err := database.DB.Order("created_at DESC").Find(&flowers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query flowers"})
	}
	c.JSON(http.StatusOK, flowers)
}

// Get /api/flowers/:id
func GetFlower(c *gin.Context) {
	id := c.Param("id")
	var flower models.Post
	if err := database.DB.First(&flower, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Flower not found"})
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

	flower := models.Post{
		Title:    payload.Title,
		Content:  payload.Content,
		ImageURL: payload.ImageURL,
		AuthorID: userId,
	}

	if err := database.DB.Create(&flower).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create flower"})
		return
	}
	c.JSON(http.StatusCreated, flower)
}

// Put /api/flowers/:id
func UpdateFlower(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var payload struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		ImageURL string `json:"image_url"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	var flower models.Post
	if err := database.DB.First(&flower, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flower not found"})
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

	isOwner := flower.AuthorID == userId
	isAdmin := false
	if rs, ok := role.(string); ok && rs == "admin" {
		isAdmin = true
	}

	if !isOwner && !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	flower.Title = payload.Title
	flower.Content = payload.Content
	flower.ImageURL = payload.ImageURL

	if err := database.DB.Save(&flower).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update flower"})
		return
	}

	c.JSON(http.StatusOK, flower)
}

// Delete /api/flowers/:id
func DeleteFlower(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var flower models.Post
	if err := database.DB.First(&flower, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flower not found"})
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

	isOwner := flower.AuthorID == userId
	isAdmin := false
	if rs, ok := role.(string); ok && rs == "admin" {
		isAdmin = true
	}

	if !isOwner && !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	if err := database.DB.Delete(&flower).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete flower"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flower deleted successfully"})
}

func LikeFlower(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var flower models.Post
	if err := database.DB.First(&flower, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flower not found"})
		return
	}

	flower.Likes++
	if err := database.DB.Save(&flower).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like flower"})
		return
	}

	c.JSON(http.StatusOK, flower)
}

func UnlikeFlower(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var flower models.Post
	if err := database.DB.First(&flower, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flower not found"})
		return
	}

	flower.Likes--
	if err := database.DB.Save(&flower).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike flower"})
		return
	}

	c.JSON(http.StatusOK, flower)
}
