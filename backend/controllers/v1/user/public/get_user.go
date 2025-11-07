package public_user_controller

import (
	user_service_factory "flower-backend/controllers/v1/user"
	publicuserdto "flower-backend/dto/public"
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
	svc, err := user_service_factory.GetUserService()
	if err != nil {
		zap.L().Error(user_service_factory.LogErrInitUserService, zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": user_service_factory.RespErrInternalServer})
		return
	}
	userId := c.Param("id")
	userIdUint, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		zap.L().Error("failed to parse user id", zap.String("user_id", userId), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	user, err := svc.GetUserByID(uint(userIdUint))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error(user_service_factory.LogErrUserNotFound, zap.String("user_id", userId))
			c.JSON(http.StatusNotFound, gin.H{"error": user_service_factory.RespErrUserNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": publicuserdto.ToPublicUser(user)})
	zap.L().Info(user_service_factory.LogMsgUserFetched, zap.Uint("user_id", uint(userIdUint)))
}

// GET /api/v1/user/:email
func GetUserByEmail(c *gin.Context) {
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
		if err == gorm.ErrRecordNotFound {
			zap.L().Error(user_service_factory.LogErrUserNotFound, zap.String("email", email))
			c.JSON(http.StatusNotFound, gin.H{"error": user_service_factory.RespErrUserNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": publicuserdto.ToPublicUser(user)})
	zap.L().Info(user_service_factory.LogMsgUserFetched, zap.String("email", email))
}

// GET /api/v1/user/:username
func GetUserByUsername(c *gin.Context) {
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
		if err == gorm.ErrRecordNotFound {
			zap.L().Error(user_service_factory.LogErrUserNotFound, zap.String("username", username))
			c.JSON(http.StatusNotFound, gin.H{"error": user_service_factory.RespErrUserNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": publicuserdto.ToPublicUser(user)})
	zap.L().Info(user_service_factory.LogMsgUserFetched, zap.String("username", username))
}

// GET /api/v1/user/all
func GetUserAll(c *gin.Context) {
	svc, err := user_service_factory.GetUserService()
	if err != nil {
		zap.L().Error(user_service_factory.LogErrInitUserService, zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": user_service_factory.RespErrInternalServer})
		return
	}

	users, err := svc.GetUserAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": publicuserdto.ToPublicUsers(users)})
	zap.L().Info("all users fetched successfully", zap.Int("users_count", len(users)))
}

// GET /api/v1/user/id/:id/select?select=field1,field2,field3
func GetUserByIDWithSelect(c *gin.Context) {
	svc, err := user_service_factory.GetUserService()
	if err != nil {
		zap.L().Error(user_service_factory.LogErrInitUserService, zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": user_service_factory.RespErrInternalServer})
		return
	}
	userId := c.Param("id")
	userIdUint, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		zap.L().Error("failed to parse user id", zap.String("user_id", userId), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	selectFieldsString := c.Query("select")
	if selectFieldsString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": user_service_factory.RespErrSelectRequired})
		return
	}
	selectFields := publicuserdto.EnsurePublicUserSelectFields(strings.Split(selectFieldsString, ","))
	user, err := svc.GetUserByIDWithSelect(uint(userIdUint), selectFields)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error(user_service_factory.LogErrUserNotFound, zap.Uint("user_id", uint(userIdUint)))
			c.JSON(http.StatusNotFound, gin.H{"error": user_service_factory.RespErrUserNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": publicuserdto.ToPublicUser(user)})
	zap.L().Info(user_service_factory.LogMsgUserFetched, zap.Uint("user_id", uint(userIdUint)))
}

// GET /api/v1/user/email/:email/select?select=field1,field2,field3
func GetUserByEmailWithSelect(c *gin.Context) {
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
	selectFieldsString := c.Query("select")
	if selectFieldsString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": user_service_factory.RespErrSelectRequired})
		return
	}
	selectFields := publicuserdto.EnsurePublicUserSelectFields(strings.Split(selectFieldsString, ","))
	user, err := svc.GetUserByEmailWithSelect(email, selectFields)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error(user_service_factory.LogErrUserNotFound, zap.String("email", email))
			c.JSON(http.StatusNotFound, gin.H{"error": user_service_factory.RespErrUserNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": publicuserdto.ToPublicUser(user)})
	zap.L().Info(user_service_factory.LogMsgUserFetched, zap.String("email", email))
}

// GET /api/v1/user/username/:username/select?select=field1,field2,field3
func GetUserByUsernameWithSelect(c *gin.Context) {
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
	selectFieldsString := c.Query("select")
	if selectFieldsString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": user_service_factory.RespErrSelectRequired})
		return
	}
	selectFields := publicuserdto.EnsurePublicUserSelectFields(strings.Split(selectFieldsString, ","))
	user, err := svc.GetUserByUsernameWithSelect(username, selectFields)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error(user_service_factory.LogErrUserNotFound, zap.String("username", username))
			c.JSON(http.StatusNotFound, gin.H{"error": user_service_factory.RespErrUserNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": publicuserdto.ToPublicUser(user)})
	zap.L().Info(user_service_factory.LogMsgUserFetched, zap.String("username", username))
}
