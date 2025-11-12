package post_controller

import (
	"flower-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GET /api/v1/post/:id
func (pc *postController) GetPostByID(c *gin.Context) {

	postId := c.Param("id")
	postIdUint, err := utils.ParseUint(postId, zap.L().Sugar())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post, err := pc.svc.GetPostByID(uint(postIdUint))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("post not found", zap.String("post_id", postId))
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
	zap.L().Info("post fetched successfully", zap.String("post_id", postId))
}

// GET /api/v1/post/user/:user_id/all
func (pc *postController) GetPostAllByUserID(c *gin.Context) {

	userId := c.Param("user_id")
	userIdUint, err := utils.ParseUint(userId, zap.L().Sugar())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	posts, err := pc.svc.GetPostAllByUserID(uint(userIdUint))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("posts not found", zap.String("user_id", userId))
			c.JSON(http.StatusNotFound, gin.H{"error": "Posts not found"})
			return
		}
		zap.L().Error("failed to get posts", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
	zap.L().Info("posts fetched successfully", zap.String("user_id", userId))
}

// GET /api/v1/post/all
func (pc *postController) GetPostAll(c *gin.Context) {
	posts, err := pc.svc.GetPostAll()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("posts not found", zap.Error(err))
			c.JSON(http.StatusNotFound, gin.H{"error": "Posts not found"})
			return
		}
		zap.L().Error("failed to get posts", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
	zap.L().Info("posts fetched successfully")
}

// GET /api/v1/post/search
func (pc *postController) SearchPosts(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query is required"})
		return
	}
	posts, err := pc.svc.SearchPosts(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search posts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
	zap.L().Info("posts searched successfully", zap.String("query", query))
}

// GET /api/v1/post/pagination
func (pc *postController) GetPostWithPagination(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")
	if page == "" || limit == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page and limit are required"})
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	posts, total, err := pc.svc.GetPostWithPagination(pageInt, limitInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts with pagination"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts, "total": total})
	zap.L().Info("posts fetched successfully with pagination", zap.Int("page", pageInt), zap.Int("limit", limitInt))
}
