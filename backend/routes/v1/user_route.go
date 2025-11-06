package v1_routes

import (
	user_controllers "flower-backend/controllers/v1/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		// Create routes
		user.POST("", user_controllers.CreateUser)
		// Get routes
		user.GET("/:id", user_controllers.GetUserByID)
		user.GET("/email/:email", user_controllers.GetUserByEmail)
		user.GET("/username/:username", user_controllers.GetUserByUsername)
		user.GET("/email/:email/select", user_controllers.GetUserByEmailWithSelect)
		user.GET("/username/:username/select", user_controllers.GetUserByUsernameWithSelect)
		user.GET("/all", user_controllers.GetUserAll)
		user.GET("/id/:id/select", user_controllers.GetUserByIDWithSelect)
		// Delete routes
		user.DELETE("/:id", user_controllers.DeleteUserByID)
		user.DELETE("/email/:email", user_controllers.DeleteUserByEmail)
		user.DELETE("/username/:username", user_controllers.DeleteUserByUsername)
		// Update routes
		user.PUT("/:id", user_controllers.UpdateUserByID)
		user.PUT("/email/:email", user_controllers.UpdateUserByEmail)
		user.PUT("/username/:username", user_controllers.UpdateUserByUsername)
	}
}
