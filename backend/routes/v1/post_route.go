package v1_routes

import (
	"flower-backend/config"
	post_controller "flower-backend/controllers/v1/post"
	"flower-backend/database"
	"flower-backend/log"
	"flower-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.RouterGroup) {
	cfg := config.LoadConfig()
	logger := log.InitLog().Sugar()
	postCtrl := post_controller.NewPostController(database.DB, cfg, logger)

	post := r.Group("/post")
	{
		// Public GET routes (no authentication required)
		post.GET("/:id", postCtrl.GetPostByID)
		post.GET("/user/:user_id/all", postCtrl.GetPostAllByUserID)
		post.GET("/all", postCtrl.GetPostAll)
		post.GET("/search", postCtrl.SearchPosts)
		post.GET("/pagination", postCtrl.GetPostWithPagination)
		post.GET("/:id/likes", postCtrl.GetPostLikes)
		post.GET("/user/:user_id/liked", postCtrl.GetUserLikedPosts)
	}

	// Protected routes (authentication required)
	postAuth := r.Group("/post")
	postAuth.Use(middlewares.Authenticate)
	{
		// Create routes
		postAuth.POST("", postCtrl.CreatePost)
		// Delete routes
		postAuth.DELETE("/:id", postCtrl.DeletePostByID)
		// Update routes
		postAuth.PUT("/:id", postCtrl.UpdatePostByIDWithSelect)
		// Like routes
		postAuth.POST("/:id/like", postCtrl.LikePost)
		postAuth.DELETE("/:id/dislike", postCtrl.DislikePost)
	}
}
