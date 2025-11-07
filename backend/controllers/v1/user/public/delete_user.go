package public_user_controller

import (
	user_service_factory "flower-backend/controllers/v1/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func DeleteUserByID(c *gin.Context) {
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
	if err := svc.DeleteUserByID(uint(userIdUint)); err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error(user_service_factory.LogErrUserNotFound, zap.String("user_id", userId))
			c.JSON(http.StatusNotFound, gin.H{"error": user_service_factory.RespErrUserNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	zap.L().Info("user deleted successfully", zap.String("user_id", userId))
}
