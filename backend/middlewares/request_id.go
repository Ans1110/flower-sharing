package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const requestIDKey = "request_id"
const requestIDHeader = "X-Request-ID"

// RequestID ensures every request has a stable request ID that is returned in
// the response header and attached to the context for downstream logging.
func RequestID(baseLogger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader(requestIDHeader)
		if reqID == "" {
			reqID = uuid.NewString()
		}

		// Expose on response and context
		c.Writer.Header().Set(requestIDHeader, reqID)
		c.Set(requestIDKey, reqID)

		// Use a request-scoped logger
		logger := baseLogger.With(zap.String(requestIDKey, reqID))
		c.Set("logger", logger)

		start := time.Now()
		c.Next()

		// Minimal completion log with request id
		logger.Info("request completed",
			zap.String("method", c.Request.Method),
			zap.String("path", c.FullPath()),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", time.Since(start)),
			zap.String("client_ip", c.ClientIP()),
		)
	}
}
