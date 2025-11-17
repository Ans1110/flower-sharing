package public_user_controller

import (
	publicuserdto "flower-backend/dto/public"
	"flower-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GET /api/v1/user/:id
func (uc *userController) GetUserByID(c *gin.Context) {
	userId := c.Param("id")
	userIdUint, err := utils.ParseUint(userId, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.svc.GetUserByID(uint(userIdUint))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.logger.Error("user not found", zap.String("user_id", userId))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": publicuserdto.ToPublicUser(user)})
	uc.logger.Info("user fetched successfully", zap.Uint("user_id", uint(userIdUint)))
}

// GET /api/v1/user/:email
func (uc *userController) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	if !utils.ValidateEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}
	user, err := uc.svc.GetUserByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.logger.Error("user not found", zap.String("email", email))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": publicuserdto.ToPublicUser(user)})
	uc.logger.Info("user fetched successfully", zap.String("email", email))
}

// GET /api/v1/user/:username
func (uc *userController) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	if !utils.ValidateUsername(username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
		return
	}
	user, err := uc.svc.GetUserByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.logger.Error("user not found", zap.String("username", username))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": publicuserdto.ToPublicUser(user)})
	uc.logger.Info("user fetched successfully", zap.String("username", username))
}

// GET /api/v1/user/all
func (uc *userController) GetUserAll(c *gin.Context) {
	users, err := uc.svc.GetUserAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": publicuserdto.ToPublicUsers(users)})
	uc.logger.Info("all users fetched successfully", zap.Int("users_count", len(users)))
}

// GET /api/v1/user/id/:id/select?select=field1,field2,field3
func (uc *userController) GetUserByIDWithSelect(c *gin.Context) {
	userId := c.Param("id")
	userIdUint, err := utils.ParseUint(userId, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	selectFieldsString := c.Query("select")
	if selectFieldsString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Select fields are required"})
		return
	}
	selectFields := publicuserdto.EnsurePublicUserSelectFields(strings.Split(selectFieldsString, ","))
	user, err := uc.svc.GetUserByIDWithSelect(uint(userIdUint), selectFields)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.logger.Error("user not found", zap.Uint("user_id", uint(userIdUint)))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": publicuserdto.ToPublicUser(user)})
	uc.logger.Info("user fetched successfully", zap.Uint("user_id", uint(userIdUint)))
}
