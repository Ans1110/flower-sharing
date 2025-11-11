package admin_user_controller

import (
	user_service_factory "flower-backend/controllers/v1/user"
	admin_user_dto "flower-backend/dto/admin"
	"flower-backend/middlewares"
	"flower-backend/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// PUT /api/v1/admin/user/:id
func UpdateUserByID(c *gin.Context) {
	svc, err := user_service_factory.GetUserService()
	if err != nil {
		zap.L().Error(user_service_factory.LogErrInitUserService, zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": user_service_factory.RespErrInternalServer})
		return
	}
	userID := c.Param("id")
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		zap.L().Error("failed to parse user id", zap.String("user_id", userID), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	user, err := svc.GetUserByID(uint(userIDUint))
	if err != nil {
		zap.L().Error("failed to get user by id", zap.String("user_id", userID), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		if middlewares.ExtractValidationErrors(c, err) {
			return
		}
		zap.L().Error(user_service_factory.LogErrBindUser, zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = svc.UpdateUserByID(uint(userIDUint), *user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error(user_service_factory.LogErrUserNotFound, zap.String("user_id", userID))
			c.JSON(http.StatusNotFound, gin.H{"error": user_service_factory.RespErrUserNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	zap.L().Info(user_service_factory.LogMsgUserUpdated, zap.String("user_id", userID))
}

// PUT /api/v1/admin/user/email/:email
func UpdateUserByEmail(c *gin.Context) {
	svc, err := user_service_factory.GetUserService()
	if err != nil {
		zap.L().Error(user_service_factory.LogErrInitUserService, zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": user_service_factory.RespErrInternalServer})
		return
	}
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}
	if !utils.ValidateEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}
	user, err := svc.GetUserByEmail(email)
	if err != nil {
		zap.L().Error("failed to get user by email", zap.String("email", email), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		if middlewares.ExtractValidationErrors(c, err) {
			return
		}
		zap.L().Error(user_service_factory.LogErrBindUser, zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = svc.UpdateUserByEmail(email, *user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error(user_service_factory.LogErrUserNotFound, zap.String("email", email))
			c.JSON(http.StatusNotFound, gin.H{"error": user_service_factory.RespErrUserNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	zap.L().Info(user_service_factory.LogMsgUserUpdated, zap.String("email", email))
}

// PUT /api/v1/admin/user/:username
func UpdateUserByUsername(c *gin.Context) {
	svc, err := user_service_factory.GetUserService()
	if err != nil {
		zap.L().Error(user_service_factory.LogErrInitUserService, zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": user_service_factory.RespErrInternalServer})
		return
	}
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}
	if !utils.ValidateUsername(username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
		return
	}
	user, err := svc.GetUserByUsername(username)
	if err != nil {
		zap.L().Error("failed to get user by username", zap.String("username", username), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		if middlewares.ExtractValidationErrors(c, err) {
			return
		}
		zap.L().Error(user_service_factory.LogErrBindUser, zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = svc.UpdateUserByUsername(username, *user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error(user_service_factory.LogErrUserNotFound, zap.String("username", username))
			c.JSON(http.StatusNotFound, gin.H{"error": user_service_factory.RespErrUserNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	zap.L().Info(user_service_factory.LogMsgUserUpdated, zap.String("username", username))
}

// PUT /api/v1/admin/user/id/:id/select?select=field1,field2,field3
func UpdateUserByIDWithSelect(c *gin.Context) {
	svc, err := user_service_factory.GetUserService()
	if err != nil {
		zap.L().Error(user_service_factory.LogErrInitUserService, zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": user_service_factory.RespErrInternalServer})
		return
	}
	userID := c.Param("id")
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		zap.L().Error("failed to parse user id", zap.String("user_id", userID), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	selectFieldsString := c.Query("select")
	if selectFieldsString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": user_service_factory.RespErrSelectRequired})
		return
	}
	selectFields := admin_user_dto.EnsureUserAdminSelectFields(strings.Split(selectFieldsString, ","))
	user, err := svc.GetUserByIDWithSelect(uint(userIDUint), selectFields)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error(user_service_factory.LogErrUserNotFound, zap.Uint("user_id", uint(userIDUint)))
			c.JSON(http.StatusNotFound, gin.H{"error": user_service_factory.RespErrUserNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		if middlewares.ExtractValidationErrors(c, err) {
			return
		}
		zap.L().Error(user_service_factory.LogErrBindUser, zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = svc.UpdateUserByIDWithSelect(uint(userIDUint), *user, selectFields)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error(user_service_factory.LogErrUserNotFound, zap.Uint("user_id", uint(userIDUint)))
			c.JSON(http.StatusNotFound, gin.H{"error": user_service_factory.RespErrUserNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": admin_user_dto.ToUserAdminDTO(user)})
	zap.L().Info(user_service_factory.LogMsgUserUpdated, zap.Uint("user_id", uint(userIDUint)))
}
