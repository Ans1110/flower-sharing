package auth_controller

import (
	"flower-backend/database"
	"flower-backend/libs"
	"flower-backend/models"
	"flower-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Avatar   string `json:"avatar"`
}

// Register godoc
//
//	@Summary		Register a new user
//	@Description	Create a new user account with username, email, password, and optional avatar
//	@Tags			auth
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			username	formData	string					true	"Username (3-20 characters)"
//	@Param			email		formData	string					true	"Email address"
//	@Param			password	formData	string					true	"Password (min 8 characters, must include uppercase, lowercase, digit, and special character)"
//	@Param			avatar		formData	file					false	"Avatar image file"
//	@Success		201			{object}	map[string]interface{}	"User registered successfully"
//	@Failure		400			{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		409			{object}	map[string]interface{}	"Conflict - user already exists"
//	@Failure		500			{object}	map[string]interface{}	"Internal server error"
//	@Router			/auth/register [post]
//	@Security		BearerAuth
func (ac *authController) Register(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	avatar, err := c.FormFile("avatar")
	if err != nil {
		ac.logger.Error("failed to get avatar file", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get avatar file"})
		return
	}

	if username == "" || email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username, email and password are required"})
		return
	}

	if !utils.ValidateUsername(username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
		return
	}

	if !utils.ValidateEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	if !utils.ValidatePassword(password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		ac.logger.Error("failed to hash password", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user, err := ac.svc.RegisterUser(username, email, hashedPassword, avatar)
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
		ac.logger.Error("failed to create token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	// set cookies
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("refreshToken", refreshToken, 7*24*60*60, "/", "", ac.cfg.GO_ENV == "production", true)

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully", "user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
		"accessToken": accessToken,
	})
}
