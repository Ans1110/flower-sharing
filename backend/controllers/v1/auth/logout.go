package auth_controller

import (
	"flower-backend/database"
	"flower-backend/libs"
	"flower-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (ac *authController) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		zap.L().Error("failed to get refresh token", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is required"})
		return
	}

	userId, err := libs.VerifyRefreshToken(refreshToken)
	if err != nil {
		zap.L().Error("failed to verify refresh token", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	if err := database.DB.Where("token = ?", refreshToken).Delete(&models.Token{}).Error; err != nil {
		zap.L().Error("failed to logout", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.SetCookie("refreshToken", "", -1, "/", "", ac.cfg.GO_ENV == "production", true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully", "userId": userId})

	zap.L().Info("Logged out successfully", zap.Uint("user_id", userId))
}
