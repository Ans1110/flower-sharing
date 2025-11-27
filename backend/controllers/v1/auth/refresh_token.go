package auth_controller

import (
	"errors"
	"flower-backend/database"
	"flower-backend/libs"
	"flower-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// RefreshToken godoc
//
//	@Summary		Refresh access token
//	@Description	Get new access token using refresh token
//	@Tags			auth
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Failure		401	{object}	map[string]interface{}
//	@Security		BearerAuth
//	@Router			/auth/refresh-token [post]
func (ac *authController) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    "AuthenticationError",
			"message": "Invalid refresh token",
		})
		return
	}

	// Check if token exists in database
	var token models.Token
	if err := database.DB.Where("token = ?", refreshToken).First(&token).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    "AuthenticationError",
			"message": "Invalid refresh token",
		})
		return
	}

	// Verify refresh token
	userId, err := libs.VerifyRefreshToken(refreshToken)
	if err != nil {
		// Check if it's a token expiration error
		if errors.Is(err, jwt.ErrTokenExpired) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    "AuthenticationError",
				"message": "Refresh token expired, please login again",
			})
			return
		}

		// Invalid token error
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    "AuthenticationError",
			"message": "Invalid refresh token",
		})
		return
	}

	// Generate new access token
	accessToken := libs.GenerateAccessToken(userId)
	if accessToken == "" {
		ac.logger.Error("Error during refresh token: failed to generate access token")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "ServerError",
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Access token refreshed successfully",
		"accessToken": accessToken,
	})
}
