package post_controller

import (
	public_dto "flower-backend/dto/public"
	"flower-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LikePost godoc
//
//	@Summary		Like a post
//	@Description	Add a like to a post
//	@Tags			posts
//	@Produce		json
//	@Param			id	path		int						true	"Post ID"
//	@Success		200	{object}	map[string]interface{}	"Post liked successfully"
//	@Failure		400	{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		500	{object}	map[string]interface{}	"Internal server error"
//	@Security		BearerAuth
//	@Router			/post/{id}/like [post]
func (pc *postController) LikePost(c *gin.Context) {

	postId := c.Param("id")
	postIdUint, err := utils.ParseUint(postId, pc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := c.GetUint("user_id")
	if userId == 0 {
		pc.logger.Error("user_id not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := pc.svc.LikePost(uint(postIdUint), userId); err != nil {
		// Handle specific error cases
		if err.Error() == "post already liked" {
			// This is expected validation, log as info instead of error
			pc.logger.Info("post already liked",
				zap.Uint("post_id", uint(postIdUint)),
				zap.Uint("user_id", userId))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Log actual errors
		pc.logger.Error("failed to like post",
			zap.Uint("post_id", uint(postIdUint)),
			zap.Uint("user_id", userId),
			zap.Error(err))

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post liked successfully"})
	pc.logger.Info("post liked successfully", zap.Uint("post_id", uint(postIdUint)))
}

// DislikePost godoc
//
//	@Summary		Unlike a post
//	@Description	Remove a like from a post
//	@Tags			posts
//	@Produce		json
//	@Param			id	path		int						true	"Post ID"
//	@Success		200	{object}	map[string]interface{}	"Post disliked successfully"
//	@Failure		400	{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		500	{object}	map[string]interface{}	"Internal server error"
//	@Security		BearerAuth
//	@Router			/post/{id}/dislike [delete]
func (pc *postController) DislikePost(c *gin.Context) {

	postId := c.Param("id")
	postIdUint, err := utils.ParseUint(postId, pc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := c.GetUint("user_id")
	if err := pc.svc.DislikePost(uint(postIdUint), userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to dislike post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post disliked successfully"})
	pc.logger.Info("post disliked successfully", zap.Uint("post_id", uint(postIdUint)))
}

// GetPostLikes godoc
//
//	@Summary		Get post likes
//	@Description	Retrieve the number of likes for a post
//	@Tags			posts
//	@Produce		json
//	@Param			id	path		int						true	"Post ID"
//	@Success		200	{object}	map[string]interface{}	"Post likes fetched successfully"
//	@Failure		400	{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		500	{object}	map[string]interface{}	"Internal server error"
//	@Securuty		BearerAuth
//	@Router			/post/{id}/likes [get]
func (pc *postController) GetPostLikes(c *gin.Context) {
	postId := c.Param("id")
	postIdUint, err := utils.ParseUint(postId, pc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	likes, err := pc.svc.GetPostLikes(uint(postIdUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post likes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"likes": likes})
	pc.logger.Info("post likes fetched successfully", zap.Uint("post_id", uint(postIdUint)))
}

// GetUserLikedPosts godoc
//
//	@Summary		Get user liked posts
//	@Description	Retrieve the posts liked by a user
//	@Tags			posts
//	@Produce		json
//	@Param			user_id	path		int						true	"User ID"
//	@Success		200		{object}	map[string]interface{}	"User liked posts fetched successfully"
//	@Failure		400		{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		500		{object}	map[string]interface{}	"Internal server error"
//	@Securuty		BearerAuth
//	@Router			/post/user/{user_id}/liked [get]
func (pc *postController) GetUserLikedPosts(c *gin.Context) {
	userId := c.Param("user_id")
	userIdUint, err := utils.ParseUint(userId, pc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page := c.Query("page")
	limit := c.Query("limit")
	pageUint, err := utils.ParseUint(page, pc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limitUint, err := utils.ParseUint(limit, pc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	posts, total, err := pc.svc.GetUserLikedPosts(uint(userIdUint), int(pageUint), int(limitUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user liked posts"})
		return
	}
	postsDTO := public_dto.ToPublicPosts(posts)
	c.JSON(http.StatusOK, gin.H{"posts": postsDTO, "total": total})
	pc.logger.Info("user liked posts fetched successfully", zap.Uint("user_id", uint(userIdUint)))
}
