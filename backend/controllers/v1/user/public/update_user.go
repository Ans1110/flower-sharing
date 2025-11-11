package public_user_controller

import (
	user_service_factory "flower-backend/controllers/v1/user"
	public_user_dto "flower-backend/dto/public"
	"flower-backend/middlewares"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// PUT /api/v1/user/id/:id
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
			zap.L().Error(user_service_factory.LogErrUserNotFound, zap.Uint("user_id", uint(userIDUint)))
			c.JSON(http.StatusNotFound, gin.H{"error": user_service_factory.RespErrUserNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	zap.L().Info(user_service_factory.LogMsgUserUpdated, zap.Uint("user_id", uint(userIDUint)))
}

// PUT /api/v1/user/id/:id/select?select=field1,field2,field3
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
	selectFields := public_user_dto.EnsurePublicUserSelectFields(strings.Split(selectFieldsString, ","))
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
	c.JSON(http.StatusOK, gin.H{"user": public_user_dto.ToPublicUser(user)})
	zap.L().Info(user_service_factory.LogMsgUserUpdated, zap.Uint("user_id", uint(userIDUint)))
}
