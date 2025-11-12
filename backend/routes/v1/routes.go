package v1_routes

import (
	"flower-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	// API v1 routes
	api := r.Group("/api/v1")
	api.Use(middlewares.ValidationError)
	{
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
