package v1_routes

import (
	admin_user_controller "flower-backend/controllers/v1/user/admin"
	"flower-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.RouterGroup) {
	admin := r.Group("/admin")
	admin.Use(middlewares.Authorize([]string{"admin"}))
	{

		//user routes
		adminUser := admin.Group("/user")
		{
			adminUser.GET("/:id", admin_user_controller.GetUserByID)
			adminUser.GET("/email/:email", admin_user_controller.GetUserByEmail)
			adminUser.GET("/username/:username", admin_user_controller.GetUserByUsername)
			adminUser.GET("/all", admin_user_controller.GetUserAll)
			adminUser.GET("/id/:id/select", admin_user_controller.GetUserByIDWithSelect)
			adminUser.GET("/email/:email/select", admin_user_controller.GetUserByEmailWithSelect)
			adminUser.GET("/username/:username/select", admin_user_controller.GetUserByUsernameWithSelect)
			// Update routes
			adminUser.PUT("/:id", admin_user_controller.UpdateUserByID)
			adminUser.PUT("/email/:email", admin_user_controller.UpdateUserByEmail)
			adminUser.PUT("/username/:username", admin_user_controller.UpdateUserByUsername)
			adminUser.PUT("/id/:id/select", admin_user_controller.UpdateUserByIDWithSelect)
			// Delete routes
			adminUser.DELETE("/:id", admin_user_controller.DeleteUserByID)
			adminUser.DELETE("/email/:email", admin_user_controller.DeleteUserByEmail)
			adminUser.DELETE("/username/:username", admin_user_controller.DeleteUserByUsername)
		}
	}
}
