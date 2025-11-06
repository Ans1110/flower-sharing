package v1

import (
	authController "flower-backend/controllers/v1/auth"
	"flower-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.Use(middlewares.ValidationError)
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		auth.POST("/logout", authController.Logout)
		auth.POST("/refresh-token", authController.RefreshToken)
	}
}
