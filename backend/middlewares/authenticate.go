package middlewares

import (
	"flower-backend/libs"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Authenticate(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
		c.Abort()
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	userId, err := libs.VerifyAccessToken(token)
	if err != nil {
		// Check if it's a JWT validation error
		if ve, ok := err.(*jwt.ValidationError); ok {
			// Check if it's a token expiration error
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    "AuthenticationError",
					"message": "Access token expired, request a new one with refresh token",
				})
				c.Abort()
				return
			}

			// Handle invalid token error
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    "AuthenticationError",
				"message": "Access token invalid",
			})
			c.Abort()
			return
		}

		// Catch-all for other errors
		zap.L().Error("Error during authentication", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "ServerError",
			"message": "Internal server error",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	c.Set("userId", userId)
	zap.L().Info("User authenticated", zap.Uint("user_id", userId))
	c.Next()
}
