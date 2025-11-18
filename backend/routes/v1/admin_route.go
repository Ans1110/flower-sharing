package v1_routes

import (
	"flower-backend/config"
	admin_user_controller "flower-backend/controllers/v1/user/admin"
	"flower-backend/database"
	"flower-backend/log"
	"flower-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.RouterGroup) {
	cfg := config.LoadConfig()
	logger := log.InitLog().Sugar()
	userCtrl := admin_user_controller.NewAdminUserController(database.DB, cfg, logger)

	admin := r.Group("/admin")
	admin.Use(middlewares.Authenticate)
	admin.Use(middlewares.Authorize([]string{"admin"}))
	{

		//user routes
		adminUser := admin.Group("/user")
		{
			adminUser.GET("/:id", userCtrl.GetUserByID)
			adminUser.GET("/email/:email", userCtrl.GetUserByEmail)
			adminUser.GET("/username/:username", userCtrl.GetUserByUsername)
			adminUser.GET("/all", userCtrl.GetUserAll)
			adminUser.GET("/id/:id/select", userCtrl.GetUserByIDWithSelect)
			// Update routes
			adminUser.PUT("/id/:id/select", userCtrl.UpdateUserByIDWithSelect)
			// Delete routes
			adminUser.DELETE("/:id", userCtrl.DeleteUserByID)
		}
	}
}
