package v1_routes

import (
	public_user_controller "flower-backend/controllers/v1/user/public"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")

	{
		// Get routes
		user.GET("/:id", public_user_controller.GetUserByID)
		user.GET("/email/:email", public_user_controller.GetUserByEmail)
		user.GET("/username/:username", public_user_controller.GetUserByUsername)
		user.GET("/email/:email/select", public_user_controller.GetUserByEmailWithSelect)
		user.GET("/username/:username/select", public_user_controller.GetUserByUsernameWithSelect)
		user.GET("/all", public_user_controller.GetUserAll)
		user.GET("/id/:id/select", public_user_controller.GetUserByIDWithSelect)
		// Delete routes
		user.DELETE("/:id", public_user_controller.DeleteUserByID)
	}
}
