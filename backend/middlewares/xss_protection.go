package middlewares

import (
	"bytes"
	"encoding/json"
	"flower-backend/utils"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// XSSProtection middleware sanitizes request body to prevent XSS attacks
func XSSProtection(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Only process for JSON content type
		contentType := c.GetHeader("Content-Type")

		if contentType == "application/json" && c.Request.Body != nil {
			// Read the body
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {
				logger.Error("failed to read request body", zap.Error(err))
				c.Next()
				return
			}

			// Restore the body for downstream handlers
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			// Parse JSON
			var data map[string]any
			if err := json.Unmarshal(bodyBytes, &data); err != nil {
				// If not valid JSON, just continue
				c.Next()
				return
			}

			// Check for XSS patterns
			if containsXSS(data) {
				logger.Warn("XSS attack attempt detected",
					zap.String("ip", c.ClientIP()),
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
				)
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": "Invalid input: potentially malicious content detected",
				})
				return
			}

			// Sanitize the data
			sanitized := utils.SanitizeMap(data)

			// Marshal back to JSON
			sanitizedBytes, err := json.Marshal(sanitized)
			if err != nil {
				logger.Error("failed to marshal sanitized data", zap.Error(err))
				c.Next()
				return
			}

			// Replace body with sanitized version
			c.Request.Body = io.NopCloser(bytes.NewBuffer(sanitizedBytes))
			c.Request.ContentLength = int64(len(sanitizedBytes))
		}

		c.Next()
	}
}

// containsXSS recursively checks for XSS patterns in data
func containsXSS(data map[string]any) bool {
	for _, value := range data {
		if hasXSSInValue(value) {
			return true
		}
	}
	return false
}

// hasXSSInValue checks for XSS patterns in a single value
func hasXSSInValue(value any) bool {
	switch v := value.(type) {
	case string:
		return utils.DetectXSSPatterns(v)
	case map[string]any:
		return containsXSS(v)
	case []any:
		return hasXSSInArray(v)
	}
	return false
}

// hasXSSInArray checks for XSS patterns in an array
func hasXSSInArray(arr []any) bool {
	for _, item := range arr {
		if hasXSSInValue(item) {
			return true
		}
	}
	return false
}

// ValidateFormInput validates and sanitizes form input fields
func ValidateFormInput() gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.GetHeader("Content-Type")

		// Skip validation for multipart/form-data as it contains file uploads
		// and ParseForm() can interfere with multipart parsing
		if len(contentType) >= 19 && contentType[:19] == "multipart/form-data" {
			c.Next()
			return
		}

		// Check all form values for XSS patterns
		if err := c.Request.ParseForm(); err == nil {
			for key, values := range c.Request.Form {
				for _, value := range values {
					if utils.DetectXSSPatterns(value) {
						c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
							"error": "Invalid input: potentially malicious content detected in field: " + key,
						})
						return
					}
				}
			}
		}

		c.Next()
	}
}
