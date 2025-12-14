package auth_controller

import (
	"flower-backend/database"
	"flower-backend/libs"
	"flower-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// Login godoc
//
//	@Summary		Login user
//	@Description	Authenticate user with email and password, returns access and refresh tokens
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		LoginRequest			true	"Login credentials"
//	@Success		200			{object}	map[string]interface{}	"Login successful, returns tokens"
//	@Failure		400			{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		401			{object}	map[string]interface{}	"Unauthorized - invalid credentials"
//	@Failure		500			{object}	map[string]interface{}	"Internal server error"
//	@Security		BearerAuth
//	@Router			/auth/login [post]
func (ac *authController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	email := req.Email
	password := req.Password

	user, err := ac.svc.GetUserByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ac.logger.Error("user not found", zap.String("email", email))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		ac.logger.Error("failed to get user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// generate access token
	accessToken := libs.GenerateAccessToken(user.ID)
	refreshToken := libs.GenerateRefreshToken(user.ID)
	expiresAt := time.Now().Add(ac.cfg.JWTRefreshExpiry)

	// create token
	database.DB.Where("user_id = ?", user.ID).Delete(&models.Token{})
	token := models.Token{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: expiresAt,
	}

	if err := database.DB.Create(&token).Error; err != nil {
		ac.logger.Error("failed to create token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	// set cookies
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("refreshToken", refreshToken, 7*24*60*60, "/", "", ac.cfg.GO_ENV == "production", true)
	c.SetCookie("role", user.Role, 7*24*60*60, "/", "", ac.cfg.GO_ENV == "production", true)

	c.JSON(http.StatusOK, gin.H{
		"message":     "Login successful",
		"accessToken": accessToken,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
			"avatar":   user.Avatar,
		},
	})

	ac.logger.Info("Login successful", zap.Any("user", user))
}
