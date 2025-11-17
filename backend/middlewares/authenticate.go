package middlewares

import (
	"errors"
	"flower-backend/libs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
		// Check if it's a token expiration error
		if errors.Is(err, jwt.ErrTokenExpired) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    "AuthenticationError",
				"message": "Access token expired, request a new one with refresh token",
			})
			c.Abort()
			return
		}

		// Check if it's a general validation error (invalid token)
		if errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, jwt.ErrTokenSignatureInvalid) {
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
