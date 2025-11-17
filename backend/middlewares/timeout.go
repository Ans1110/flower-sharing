package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	DefaultTimeoutMessage = "Request timeout - the server took too long to process your request"
	SwaggerIndexPath      = "/swagger/index.html"
	SwaggerDocPath        = "/swagger/doc.json"
	SwaggerPrefix         = "/swagger"
)

type TimeoutConfig struct {
	Timeout      time.Duration
	Logger       *zap.Logger
	ErrorMessage string
}

func DefaultTimeoutConfig(logger *zap.Logger) TimeoutConfig {
	return TimeoutConfig{
		Timeout:      30 * time.Second,
		Logger:       logger,
		ErrorMessage: DefaultTimeoutMessage,
	}
}

func isSwaggerPath(path string) bool {
	return path == SwaggerIndexPath ||
		path == SwaggerDocPath ||
		len(path) > 8 && path[:8] == SwaggerPrefix
}

// RequestTimeout creates a middleware that sets a timeout for each request
func RequestTimeout(config TimeoutConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if isSwaggerPath(c.Request.URL.Path) {
			c.Next()
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), config.Timeout)
		defer cancel()

		// Replace request context with timeout context
		c.Request = c.Request.WithContext(ctx)

		// Channel to signal when the request is done
		done := make(chan struct{})

		go func() {
			c.Next()
			close(done)
		}()

		// Wait for either the request to complete or timeout
		select {
		case <-done:
			return
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				config.Logger.Warn("request timeout",
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
					zap.String("ip", c.ClientIP()),
					zap.Duration("timeout", config.Timeout),
				)

				c.AbortWithStatusJSON(http.StatusRequestTimeout, gin.H{
					"error":   config.ErrorMessage,
					"code":    http.StatusRequestTimeout,
					"timeout": config.Timeout.String(),
				})
			}
		}
	}
}

// RequestTimeoutWithCustomDuration creates a timeout middleware with custom duration
func RequestTimeoutWithCustomDuration(timeout time.Duration, logger *zap.Logger) gin.HandlerFunc {
	config := TimeoutConfig{
		Timeout:      timeout,
		Logger:       logger,
		ErrorMessage: DefaultTimeoutMessage,
	}
	return RequestTimeout(config)
}

// ContextTimeoutMiddleware adds timeout context to requests (lighter approach)
// This doesn't abort requests but makes the context timeout available to handlers
func ContextTimeoutMiddleware(timeout time.Duration, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if isSwaggerPath(c.Request.URL.Path) {
			c.Next()
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Replace request context with timeout context
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		if ctx.Err() == context.DeadlineExceeded {
			logger.Warn("request context deadline exceeded",
				zap.String("method", c.Request.Method),
				zap.String("path", c.Request.URL.Path),
				zap.String("ip", c.ClientIP()),
				zap.Duration("timeout", timeout),
			)
		}
	}
}

// TimeoutByRoute returns different timeout durations based on route patterns
func TimeoutByRoute(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if isSwaggerPath(c.Request.URL.Path) {
			c.Next()
			return
		}

		var timeout time.Duration

		// Set different timeouts based on route patterns
		switch {
		case c.Request.Method == "GET":
			timeout = 30 * time.Second
		case c.Request.Method == "POST" && c.ContentType() == "multipart/form-data":
			timeout = 120 * time.Second
		case c.Request.Method == "PUT" && c.ContentType() == "multipart/form-data":
			timeout = 120 * time.Second
		case c.Request.Method == "POST" || c.Request.Method == "PUT":
			timeout = 60 * time.Second
		case c.Request.Method == "DELETE":
			timeout = 45 * time.Second
		default:
			timeout = 30 * time.Second
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Replace request context with timeout context
		c.Request = c.Request.WithContext(ctx)

		// Channel to signal when the request is done
		done := make(chan struct{})

		go func() {
			c.Next()
			close(done)
		}()

		// Wait for either the request to complete or timeout
		select {
		case <-done:
			return
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				logger.Warn("request timeout",
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
					zap.String("ip", c.ClientIP()),
					zap.Duration("timeout", timeout),
					zap.String("content_type", c.ContentType()),
				)

				// Abort the request and return timeout error
				c.AbortWithStatusJSON(http.StatusRequestTimeout, gin.H{
					"error":   DefaultTimeoutMessage,
					"code":    http.StatusRequestTimeout,
					"timeout": timeout.String(),
				})
			}
		}
	}
}
