package post_controller

import (
	"flower-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DeletePostByID godoc
//
//	@Summary		Delete post
//	@Description	Delete a post by ID
//	@Tags			posts
//	@Produce		json
//	@Param			id	path		int						true	"Post ID"
//	@Success		200	{object}	map[string]interface{}	"Post deleted successfully"
//	@Failure		400	{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		403	{object}	map[string]interface{}
//	@Failure		404	{object}	map[string]interface{}	"Post not found"
//	@Failure		500	{object}	map[string]interface{}	"Internal server error"
//	@Security		BearerAuth
//	@Router			/post/{id} [delete]
func (pc *postController) DeletePostByID(c *gin.Context) {

	postId := c.Param("id")
	postIdUint, err := utils.ParseUint(postId, pc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := c.GetUint("user_id")
	ownership, err := pc.svc.CheckPostOwnership(uint(postIdUint), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check post ownership"})
		return
	}
	if !ownership {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the owner of this post"})
		return
	}
	if err := pc.svc.DeletePostByID(uint(postIdUint), userId); err != nil {
		if err == gorm.ErrRecordNotFound {
			pc.logger.Error("post not found", zap.String("post_id", postId))
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
	pc.logger.Info("post deleted successfully", zap.String("post_id", postId), zap.Uint("user_id", userId))
}
