package auth_controllers

import (
	"flower-backend/config"
	"flower-backend/database"
	"flower-backend/libs"
	"flower-backend/middlewares"
	"flower-backend/models"
	user_services "flower-backend/services/v1/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var cfg = config.LoadConfig()
var userService = user_services.NewUserService(database.DB, cfg)

type RegisterInput struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func Register(c *gin.Context) {
	var body RegisterInput

	if err := c.ShouldBindJSON(&body); err != nil {
		if middlewares.ExtractValidationErrors(c, err) {
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("failed to hash password", zap.Error(err))
		return
	}

	user, err := userService.CreateUser(models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// generate refresh token
	accessToken := libs.GenerateAccessToken(user.ID)
	refreshToken := libs.GenerateRefreshToken(user.ID)

	// create token
	token := models.Token{
		Token:  refreshToken,
		UserID: user.ID,
	}

	if err := database.DB.Create(&token).Error; err != nil {
		zap.L().Error("failed to create token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	// set cookies
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("refreshToken", refreshToken, 7*24*60*60, "/", "", cfg.GO_ENV == "production", true)

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully", "user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
		"accessToken": accessToken,
	})
}
