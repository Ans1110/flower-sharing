package authController

import (
	"flower-backend/config"
	"flower-backend/database"
	"flower-backend/libs"
	"flower-backend/middlewares"
	"flower-backend/models"
	"flower-backend/services/v1"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var body LoginInput
	cfg := config.LoadConfig()
	userService := services.NewUserService(database.DB, cfg)

	if err := c.ShouldBindJSON(&body); err != nil {
		if middlewares.ExtractValidationErrors(c, err) {
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userService.GetUserByEmail(body.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("user not found", zap.String("email", body.Email))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		zap.L().Error("failed to get user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// generate access token
	accessToken := libs.GenerateAccessToken(user.ID)
	refreshToken := libs.GenerateRefreshToken(user.ID)

	// create token
	database.DB.Where("user_id = ?", user.ID).Delete(&models.Token{})
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
		"message": "Login successful",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
		"accessToken": accessToken,
	})

	zap.L().Info("Login successful", zap.String("user", user.Username), zap.String("accessToken", accessToken))
}
