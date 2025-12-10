package middlewares

import (
	"errors"
	"flower-backend/libs"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// OptionalAuthenticate is a middleware that attempts to authenticate the user
// but does not abort the request if authentication fails. It sets user_id in
// the context if a valid token is provided, otherwise continues without it.
func OptionalAuthenticate(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		// No token provided, continue without authentication
		c.Next()
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	userId, err := libs.VerifyAccessToken(token)
	if err != nil {
		// Token is invalid or expired, but don't abort - just continue without auth
		if errors.Is(err, jwt.ErrTokenExpired) ||
			errors.Is(err, jwt.ErrTokenMalformed) ||
			errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			c.Next()
			return
		}
		// For other errors, also continue without auth
		c.Next()
		return
	}

	// Valid token - set user_id in context
	c.Set("user_id", userId)
	c.Next()
}
