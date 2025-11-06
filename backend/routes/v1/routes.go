package v1

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	// API v1 routes
	api := r.Group("/api/v1")
	{
		// Auth routes
		// /api/v1/auth
		AuthRoutes(api)
		// TODO: Add more routes here
		// User routes
		// Post routes
	}
}
