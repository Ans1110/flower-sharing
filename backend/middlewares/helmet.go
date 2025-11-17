package middlewares

import (
	"flower-backend/config"

	"github.com/gin-gonic/gin"
)

func Helmet() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.LoadConfig()

		// Prevent clickjacking attacks by disallowing embedding in iframes
		c.Header("X-Frame-Options", "DENY")

		// Enable XSS filter built into most browsers
		c.Header("X-XSS-Protection", "1; mode=block")

		// Prevent MIME type sniffing
		c.Header("X-Content-Type-Options", "nosniff")

		// Control referrer information
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")

		csp := "default-src 'self'; " +
			"script-src 'self'; " + // Remove unsafe-inline and unsafe-eval for better security
			"style-src 'self' 'unsafe-inline'; " + // Allow inline styles (can be restricted further)
			"img-src 'self' data: https: blob:; " + // Allow images from HTTPS and data URIs (for Cloudinary, etc.)
			"font-src 'self' data:; " +
			"connect-src 'self'; " +
			"frame-src 'none'; " + // Block all frames
			"frame-ancestors 'none'; " + // Block all frame ancestors
			"media-src 'self'; " +
			"object-src 'none'; " + // Block plugins like Flash
			"base-uri 'self'; " + // Prevent base tag injection
			"form-action 'self'; " + // Restrict form submissions
			"upgrade-insecure-requests; " + // Upgrade HTTP to HTTPS
			"block-all-mixed-content;" // Block mixed content

		c.Header("Content-Security-Policy", csp)

		// Enforce HTTPS in production
		if cfg.GO_ENV == "production" {
			c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		}

		// Restrict browser features
		c.Header("Permissions-Policy", "geolocation=(), midi=(), sync-xhr=(), microphone=(), camera=(), magnetometer=(), gyroscope=(), fullscreen=(self), payment=(), usb=(), interest-cohort=()")

		// Prevent DNS prefetching
		c.Header("X-DNS-Prefetch-Control", "off")

		// Prevent browser caching of sensitive data
		if c.Request.URL.Path != "/swagger/*any" {
			c.Header("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
			c.Header("Pragma", "no-cache")
			c.Header("Expires", "0")
		}

		c.Next()
	}
}
