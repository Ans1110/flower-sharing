package middlewares

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var skipRoutes string = "/api/v1"

func HttpLogger(c *gin.Context) {
	path := c.Request.URL.Path

	if strings.Contains(c.Request.URL.Path, skipRoutes) {
		c.Next()
		return
	}

	start := time.Now()
	method := c.Request.Method
	ip := c.ClientIP()
	userAgent := c.Request.UserAgent()
	statusCode := strconv.Itoa(c.Writer.Status())
	protocol := c.Request.Proto
	hostname := c.Request.Host
	duration := time.Since(start)

	if strings.Contains(c.Request.URL.Path, `/\.(jpg|jpeg|png|gif|svg|ico|woff2|css|js)$/`) {
		return
	}

	zap.L().Info("HTTP request", zap.String("path", path), zap.String("method", method), zap.String("statusCode", statusCode), zap.String("ip", ip), zap.String("userAgent", userAgent), zap.String("statusCode", statusCode), zap.String("protocol", protocol), zap.String("hostname", hostname), zap.Duration("duration", duration))
	c.Next()
}
