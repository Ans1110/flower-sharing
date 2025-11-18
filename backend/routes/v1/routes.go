package v1_routes

import (
	"flower-backend/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"message": "Flower Sharing API is running",
		})
	})

	// API v1 routes
	api := r.Group("/api/v1")
	api.Use(middlewares.ValidationError)
	{
		// Health check endpoint
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  "healthy",
				"version": "1.0",
				"message": "Flower Sharing API v1 is running",
			})
		})

		// Auth routes
		// /api/v1/auth
		AuthRoutes(api)
		// User routes
		// /api/v1/user
		UserRoutes(api)
		// Admin routes
		// /api/v1/admin
		AdminRoutes(api)
		// Post routes
		// /api/v1/post
		PostRoutes(api)
	}
}
