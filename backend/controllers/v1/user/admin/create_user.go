package admin_user_controller

import (
	user_service_factory "flower-backend/controllers/v1/user"
	"flower-backend/middlewares"
	"flower-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// POST /api/v1/admin/user
func CreateUser(c *gin.Context) {
	var body models.User

	userService, err := user_service_factory.GetUserService()
	if err != nil {
		zap.L().Error(user_service_factory.LogErrInitUserService, zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": user_service_factory.RespErrInternalServer})
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		if middlewares.ExtractValidationErrors(c, err) {
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := userService.CreateUser(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	zap.L().Info("user created successfully", zap.String("username", user.Username))
}
