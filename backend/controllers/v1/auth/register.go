package auth_controller

import (
	"flower-backend/libs"
	"flower-backend/models"
	"flower-backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=2,max=15"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// Register godoc
//
//	@Summary		Register a new user
//	@Description	Create a new user account with username, email, and password
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		RegisterRequest			true	"Registration credentials"
//	@Success		200			{object}	map[string]interface{}	"User registered successfully"
//	@Failure		400			{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		409			{object}	map[string]interface{}	"Conflict - user already exists"
//	@Failure		500			{object}	map[string]interface{}	"Internal server error"
//	@Security		BearerAuth
//	@Router			/auth/register [post]
func (ac *authController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONError(c, http.StatusBadRequest, "ValidationError", "Username, email and password are required")
		return
	}

	username := req.Username
	email := req.Email
	password := req.Password

	if !utils.ValidateUsername(username) {
		utils.JSONError(c, http.StatusBadRequest, "ValidationError", "Invalid username")
		return
	}

	if !utils.ValidateEmail(email) {
		utils.JSONError(c, http.StatusBadRequest, "ValidationError", "Invalid email")
		return
	}

	if !utils.ValidatePassword(password) {
		utils.JSONError(c, http.StatusBadRequest, "ValidationError", "Invalid password")
		return
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		ac.logger.Error("failed to hash password", zap.Error(err))
		utils.JSONError(c, http.StatusInternalServerError, "", "Failed to hash password")
		return
	}

	user, err := ac.svc.RegisterUser(username, email, hashedPassword, nil)
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "", "Failed to register user")
		return
	}

	// generate refresh token
	accessToken := libs.GenerateAccessToken(user.ID)
	refreshToken := libs.GenerateRefreshToken(user.ID)
	expiresAt := time.Now().Add(ac.cfg.JWTRefreshExpiry)

	// create token
	token := models.Token{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: expiresAt,
	}

	if err := ac.svc.CreateToken(&token); err != nil {
		ac.logger.Error("failed to create token", zap.Error(err))
		utils.JSONError(c, http.StatusInternalServerError, "", "Failed to create token")
		return
	}

	// set cookies
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("refreshToken", refreshToken, int(7*24*60*60), "/", "", ac.cfg.GO_ENV == "production", true)
	c.SetCookie("role", user.Role, 7*24*60*60, "/", "", ac.cfg.GO_ENV == "production", true)

	c.JSON(http.StatusOK, gin.H{
		"message":     "User created successfully",
		"accessToken": accessToken,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
			"avatar":   user.Avatar,
		},
	})
}
