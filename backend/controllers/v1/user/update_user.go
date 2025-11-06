package user_controllers

import (
	"flower-backend/middlewares"
	"flower-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// PUT /api/v1/user/:id
func UpdateUserByID(c *gin.Context) {
	userID := c.Param("id")
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		zap.L().Error("failed to parse user id", zap.String("user_id", userID), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	user, err := userService.GetUserByID(uint(userIDUint))
	if err != nil {
		zap.L().Error("failed to get user by id", zap.String("user_id", userID), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		if middlewares.ExtractValidationErrors(c, err) {
			return
		}
		zap.L().Error("failed to bind user", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = userService.UpdateUserByID(uint(userIDUint), *user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("user not found", zap.String("user_id", userID))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	zap.L().Info("user updated successfully", zap.String("user_id", userID))
}

// PUT /api/v1/user/email/:email
func UpdateUserByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}
	if !utils.ValidateEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}
	user, err := userService.GetUserByEmail(email)
	if err != nil {
		zap.L().Error("failed to get user by email", zap.String("email", email), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		if middlewares.ExtractValidationErrors(c, err) {
			return
		}
		zap.L().Error("failed to bind user", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = userService.UpdateUserByEmail(email, *user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("user not found", zap.String("email", email))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	zap.L().Info("user updated successfully", zap.String("email", email))
}

// PUT /api/v1/user/:username
func UpdateUserByUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}
	if !utils.ValidateUsername(username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
		return
	}
	user, err := userService.GetUserByUsername(username)
	if err != nil {
		zap.L().Error("failed to get user by username", zap.String("username", username), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		if middlewares.ExtractValidationErrors(c, err) {
			return
		}
		zap.L().Error("failed to bind user", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = userService.UpdateUserByUsername(username, *user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("user not found", zap.String("username", username))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	zap.L().Info("user updated successfully", zap.String("username", username))
}
