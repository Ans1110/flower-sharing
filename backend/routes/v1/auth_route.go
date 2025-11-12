package v1_routes

import (
	"flower-backend/config"
	auth_controller "flower-backend/controllers/v1/auth"
	"flower-backend/database"
	"flower-backend/log"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	cfg := config.LoadConfig()
	logger := log.InitLog().Sugar()
	authCtrl := auth_controller.NewAuthController(database.DB, cfg, logger)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authCtrl.Register)
		auth.POST("/login", authCtrl.Login)
		auth.POST("/logout", authCtrl.Logout)
		auth.POST("/refresh-token", authCtrl.RefreshToken)
	}
}
