package v1_routes

import (
	"flower-backend/config"
	public_user_controller "flower-backend/controllers/v1/user/public"
	"flower-backend/database"
	"flower-backend/log"
	"flower-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	cfg := config.LoadConfig()
	logger := log.InitLog().Sugar()
	userCtrl := public_user_controller.NewUserController(database.DB, cfg, logger)

	user := r.Group("/user")
	{
		// Public GET routes (no authentication required)
		user.GET("/:id", userCtrl.GetUserByID)
		user.GET("/email/:email", userCtrl.GetUserByEmail)
		user.GET("/username/:username", userCtrl.GetUserByUsername)
		user.GET("/all", userCtrl.GetUserAll)
		user.GET("/id/:id/select", userCtrl.GetUserByIDWithSelect)
		user.GET("/followers/:user_id", userCtrl.GetUserFollowers)
		user.GET("/following/:user_id", userCtrl.GetUserFollowing)
		user.GET("/followers-count/:user_id", userCtrl.GetUserFollowersCount)
		user.GET("/following-count/:user_id", userCtrl.GetUserFollowingCount)
		user.GET("/following-posts/:user_id", userCtrl.GetUserFollowingPosts)
	}

	// Protected routes (authentication required)
	userAuth := r.Group("/user")
	userAuth.Use(middlewares.Authenticate)
	{
		// Update routes
		userAuth.PUT("/id/:id/select", userCtrl.UpdateUserByIDWithSelect)
		// Delete routes
		userAuth.DELETE("/:id", userCtrl.DeleteUserByID)
		// Follow routes
		userAuth.POST("/follow/:follower_id/:following_id", userCtrl.FollowUser)
		userAuth.POST("/unfollow/:follower_id/:following_id", userCtrl.UnfollowUser)
	}
}
