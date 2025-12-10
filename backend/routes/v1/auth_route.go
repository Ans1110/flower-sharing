package v1_routes

import (
	"flower-backend/config"
	auth_controller "flower-backend/controllers/v1/auth"
	"flower-backend/database"
	"flower-backend/log"
	"flower-backend/middlewares"

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

		// OAuth routes
		auth.GET("/google", authCtrl.GoogleLogin)
		auth.GET("/google/callback", authCtrl.GoogleCallback)
		auth.GET("/github", authCtrl.GithubLogin)
		auth.GET("/github/callback", authCtrl.GithubCallback)
	}

	// Protected auth routes
	authProtected := r.Group("/auth")
	authProtected.Use(middlewares.Authenticate)
	{
		authProtected.GET("/me", authCtrl.Me)
	}
}
