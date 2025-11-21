package auth_controller

import (
	"flower-backend/database"
	"flower-backend/libs"
	"flower-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logout godoc
//
//	@Summary		Logout user
//	@Description	Logout user by invalidating the refresh token
//	@Tags			auth
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"Logged out successfully"
//	@Failure		401	{object}	map[string]interface{}	"Unauthorized - invalid or missing refresh token"
//	@Failure		500	{object}	map[string]interface{}	"Internal server error"
//	@Security		BearerAuth
//	@Router			/auth/logout [post]
func (ac *authController) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		ac.logger.Error("failed to get refresh token", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is required"})
		return
	}

	userId, err := libs.VerifyRefreshToken(refreshToken)
	if err != nil {
		ac.logger.Error("failed to verify refresh token", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	if err := database.DB.Where("token = ?", refreshToken).Delete(&models.Token{}).Error; err != nil {
		ac.logger.Error("failed to logout", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.SetCookie("refreshToken", "", -1, "/", "", ac.cfg.GO_ENV == "production", true)
	c.SetCookie("accessToken", "", -1, "/", "", ac.cfg.GO_ENV == "production", true)
	c.SetCookie("role", "", -1, "/", "", ac.cfg.GO_ENV == "production", true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully", "userId": userId})

	ac.logger.Info("Logged out successfully", zap.Uint("user_id", userId))
}
