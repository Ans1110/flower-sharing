package v1_routes

import (
	public_user_controller "flower-backend/controllers/v1/user/public"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")

	{
		// Create routes
		user.POST("", public_user_controller.CreateUser)
		// Get routes
		user.GET("/:id", public_user_controller.GetUserByID)
		user.GET("/email/:email", public_user_controller.GetUserByEmail)
		user.GET("/username/:username", public_user_controller.GetUserByUsername)
		user.GET("/email/:email/select", public_user_controller.GetUserByEmailWithSelect)
		user.GET("/username/:username/select", public_user_controller.GetUserByUsernameWithSelect)
		user.GET("/all", public_user_controller.GetUserAll)
		user.GET("/id/:id/select", public_user_controller.GetUserByIDWithSelect)
		// Update routes
		user.PUT("/:id", public_user_controller.UpdateUserByID)
		user.PUT("/:id/select", public_user_controller.UpdateUserByIDWithSelect)
		// Delete routes
		user.DELETE("/:id", public_user_controller.DeleteUserByID)
	}
}
