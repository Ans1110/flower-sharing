package user_controllers

import (
	"flower-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DELETE /api/v1/user/:id
func DeleteUserByID(c *gin.Context) {
	userId := c.Param("id")
	userIdUint, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		zap.L().Error("failed to parse user id", zap.String("user_id", userId), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	err = userService.DeleteUserByID(uint(userIdUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// DELETE /api/v1/user/email/:email
func DeleteUserByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}
	if !utils.ValidateEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}
	err := userService.DeleteUserByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("user not found", zap.String("email", email))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	zap.L().Info("user deleted successfully", zap.String("email", email))
}

// DELETE /api/v1/user/username/:username
func DeleteUserByUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}
	if !utils.ValidateUsername(username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
		return
	}
	err := userService.DeleteUserByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("user not found", zap.String("username", username))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	zap.L().Info("user deleted successfully", zap.String("username", username))
}
