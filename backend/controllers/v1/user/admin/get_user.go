package admin_user_controller

import (
	admin_user_dto "flower-backend/dto/admin"
	"flower-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GET /api/v1/admin/user/:id
func (uc *adminUserController) GetUserByID(c *gin.Context) {
	userId := c.Param("id")
	userIdUint, err := utils.ParseUint(userId, zap.L().Sugar())
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": admin_user_dto.ToUserAdminDTO(user)})
	uc.logger.Info("user fetched successfully", zap.Uint("user_id", uint(userIdUint)))
}

// GET /api/v1/admin/user/:email
func (uc *adminUserController) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	if !utils.ValidateEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}
	user, err := uc.svc.GetUserByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.logger.Error("user not found", zap.String("email", email))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": admin_user_dto.ToUserAdminDTO(user)})
	uc.logger.Info("user fetched successfully", zap.String("email", email))
}

// GET /api/v1/admin/user/:username
func (uc *adminUserController) GetUserByUsername(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{"user": admin_user_dto.ToUserAdminDTO(user)})
	uc.logger.Info("user fetched successfully", zap.String("username", username))
}

// GET /api/v1/admin/user/all
func (uc *adminUserController) GetUserAll(c *gin.Context) {
	users, err := uc.svc.GetUserAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": admin_user_dto.ToUserAdminDTOs(users)})
	zap.L().Info("all users fetched successfully", zap.Int("users_count", len(users)))
}

// GET /api/v1/admin/user/id/:id/select?select=field1,field2,field3
func (uc *adminUserController) GetUserByIDWithSelect(c *gin.Context) {
	userId := c.Param("id")
	userIdUint, err := utils.ParseUint(userId, zap.L().Sugar())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	selectFieldsString := c.Query("select")
	if selectFieldsString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Select fields are required"})
		return
	}
	selectFields := admin_user_dto.EnsureUserAdminSelectFields(strings.Split(selectFieldsString, ","))
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
	c.JSON(http.StatusOK, gin.H{"user": admin_user_dto.ToUserAdminDTO(user)})
	uc.logger.Info("user fetched successfully", zap.Uint("user_id", uint(userIdUint)))
}
