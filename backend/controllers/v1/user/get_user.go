package user_controllers

import (
	"flower-backend/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GET /api/v1/user/:id
func GetUserByID(c *gin.Context) {
	userId := c.Param("id")
	userIdUint, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		zap.L().Error("failed to parse user id", zap.String("user_id", userId), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	user, err := userService.GetUserByID(uint(userIdUint))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("user not found", zap.String("user_id", userId))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	zap.L().Info("user fetched successfully", zap.Uint("user_id", uint(userIdUint)))
}

// GET /api/v1/user/:email
func GetUserByEmail(c *gin.Context) {
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
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("user not found", zap.String("email", email))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	zap.L().Info("user fetched successfully", zap.String("email", email))
}

// GET /api/v1/user/:username
func GetUserByUsername(c *gin.Context) {
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
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("user not found", zap.String("username", username))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	zap.L().Info("user fetched successfully", zap.String("username", username))
}

// GET /api/v1/user/all
func GetUserAll(c *gin.Context) {
	users, err := userService.GetUserAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
	zap.L().Info("all users fetched successfully", zap.Int("users_count", len(users)))
}

// GET /api/v1/user/:id/select?select=field1,field2,field3
func GetUserByIDWithSelect(c *gin.Context) {
	userId := c.Param("id")
	userIdUint, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		zap.L().Error("failed to parse user id", zap.String("user_id", userId), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	selectFieldsString := c.Query("select")
	if selectFieldsString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Select fields are required"})
		return
	}
	selectFields := strings.Split(selectFieldsString, ",")
	user, err := userService.GetUserByIDWithSelect(uint(userIdUint), selectFields)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("user not found", zap.Uint("user_id", uint(userIdUint)))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	zap.L().Info("user fetched successfully", zap.Uint("user_id", uint(userIdUint)))
}

// GET /api/v1/user/:email/select?select=field1,field2,field3
func GetUserByEmailWithSelect(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}
	if !utils.ValidateEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}
	selectFieldsString := c.Query("select")
	if selectFieldsString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Select fields are required"})
		return
	}
	selectFields := strings.Split(selectFieldsString, ",")
	user, err := userService.GetUserByEmailWithSelect(email, selectFields)
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
	zap.L().Info("user fetched successfully", zap.String("email", email))
}

// GET /api/v1/user/:username/select?select=field1,field2,field3
func GetUserByUsernameWithSelect(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}
	selectFieldsString := c.Query("select")
	if selectFieldsString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Select fields are required"})
		return
	}
	selectFields := strings.Split(selectFieldsString, ",")
	user, err := userService.GetUserByUsernameWithSelect(username, selectFields)
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
	zap.L().Info("user fetched successfully", zap.String("username", username))
}
