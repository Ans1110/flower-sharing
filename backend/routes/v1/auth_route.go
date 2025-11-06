package v1_routes

import (
	auth_controllers "flower-backend/controllers/v1/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", auth_controllers.Register)
		auth.POST("/login", auth_controllers.Login)
		auth.POST("/logout", auth_controllers.Logout)
		auth.POST("/refresh-token", auth_controllers.RefreshToken)
	}
}
