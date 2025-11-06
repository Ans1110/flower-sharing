package user_controllers

import (
	"flower-backend/config"
	"flower-backend/database"
	"flower-backend/middlewares"
	"flower-backend/models"
	user_services "flower-backend/services/v1/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var userService = user_services.NewUserService(database.DB, config.LoadConfig())

// POST /api/v1/user
func CreateUser(c *gin.Context) {
	var body models.User
	cfg := config.LoadConfig()
	userService := user_services.NewUserService(database.DB, cfg)

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
