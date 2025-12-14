package middlewares

import (
	"flower-backend/config"
	"flower-backend/libs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	csrfCookieName = "csrf_token"
	csrfHeaderName = "X-CSRF-Token"
)

// CSRFProtection enforces double-submit-cookie CSRF validation for state-changing requests.
// - Issues a non-HttpOnly SameSite cookie with a random token (24h TTL) when missing.
// - Requires the same token to be sent in the X-CSRF-Token header for unsafe methods.
// - Skips safe methods: GET, HEAD, OPTIONS.
func CSRFProtection(cfg *config.Config, logger *zap.Logger) gin.HandlerFunc {
	secure := cfg.GO_ENV == "production"
	maxAge := int((24 * time.Hour).Seconds())

	return func(c *gin.Context) {
		method := c.Request.Method

		csrfToken, err := c.Cookie(csrfCookieName)
		if err != nil || csrfToken == "" {
			csrfToken = libs.GenerateRandomString(32)
			// Non-HttpOnly to allow frontend to read and echo in header; SameSite reduces cross-site leakage.
			c.SetSameSite(http.SameSiteLaxMode)
			c.SetCookie(csrfCookieName, csrfToken, maxAge, "/", "", secure, false)
			// Surface token in response header to help initial fetchers.
			c.Header(csrfHeaderName, csrfToken)
		}

		// Allow safe methods without header validation.
		if method == http.MethodGet || method == http.MethodHead || method == http.MethodOptions {
			c.Next()
			return
		}

		headerToken := c.GetHeader(csrfHeaderName)
		if headerToken == "" {
			// Fallback to common alternate header spelling.
			headerToken = c.GetHeader("X-CSRFToken")
		}

		if headerToken == "" || headerToken != csrfToken {
			logger.Warn("csrf token validation failed",
				zap.String("ip", c.ClientIP()),
				zap.String("path", c.FullPath()),
				zap.String("method", method),
			)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "csrf validation failed",
				"message": "Invalid or missing CSRF token",
			})
			return
		}

		c.Next()
	}
}
