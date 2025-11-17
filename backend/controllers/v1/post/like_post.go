package post_controller

import (
	"flower-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// POST /api/v1/post/:id/like
func (pc *postController) LikePost(c *gin.Context) {

	postId := c.Param("id")
	postIdUint, err := utils.ParseUint(postId, pc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := c.GetUint("user_id")
	if err := pc.svc.LikePost(uint(postIdUint), userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post liked successfully"})
	pc.logger.Info("post liked successfully", zap.Uint("post_id", uint(postIdUint)))
}

// DELETE /api/v1/post/:id/dislike
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

// GET /api/v1/post/:id/likes
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

// GET /api/v1/post/user/:user_id/liked
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
	c.JSON(http.StatusOK, gin.H{"posts": posts, "total": total})
	pc.logger.Info("user liked posts fetched successfully", zap.Uint("user_id", uint(userIdUint)))
}
