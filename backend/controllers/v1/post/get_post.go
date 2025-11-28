package post_controller

import (
	"flower-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GetPostByID godoc
//
//	@Summary		Get post by ID
//	@Description	Retrieve a single post by its ID
//	@Tags			posts
//	@Produce		json
//	@Param			id	path		int	true	"Post ID"
//	@Success		200	{object}	map[string]interface{}
//	@Failure		404	{object}	map[string]interface{}
//	@Securuty		BearerAuth
//	@Router			/post/{id} [get]
func (pc *postController) GetPostByID(c *gin.Context) {

	postId := c.Param("id")
	postIdUint, err := utils.ParseUint(postId, pc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post, err := pc.svc.GetPostByID(uint(postIdUint))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			pc.logger.Error("post not found", zap.String("post_id", postId))
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		pc.logger.Error("failed to get post", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
	pc.logger.Info("post fetched successfully", zap.String("post_id", postId))
}

// GetPostAllByUserID godoc
//
//	@Summary		Get all posts by user ID
//	@Description	Retrieve all posts by a specific user
//	@Tags			posts
//	@Produce		json
//	@Param			user_id	path		int						true	"User ID"
//	@Success		200		{object}	map[string]interface{}	"Posts fetched successfully"
//	@Failure		404		{object}	map[string]interface{}	"Posts not found"
//	@Failure		500		{object}	map[string]interface{}	"Internal server error"
//	@Securuty		BearerAuth
//	@Router			/post/user/{user_id}/all [get]
func (pc *postController) GetPostAllByUserID(c *gin.Context) {

	userId := c.Param("user_id")
	userIdUint, err := utils.ParseUint(userId, pc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	posts, err := pc.svc.GetPostAllByUserID(uint(userIdUint))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			pc.logger.Error("posts not found", zap.String("user_id", userId))
			c.JSON(http.StatusNotFound, gin.H{"error": "Posts not found"})
			return
		}
		pc.logger.Error("failed to get posts", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
	pc.logger.Info("posts fetched successfully", zap.String("user_id", userId))
}

// GetPostAll godoc
//
//	@Summary		Get all posts
//	@Description	Retrieve all posts
//	@Tags			posts
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"Posts fetched successfully"
//	@Failure		404	{object}	map[string]interface{}	"Posts not found"
//	@Failure		500	{object}	map[string]interface{}	"Internal server error"
//	@Securuty		BearerAuth
//	@Router			/post/all [get]
func (pc *postController) GetPostAll(c *gin.Context) {
	posts, err := pc.svc.GetPostAll()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			pc.logger.Error("posts not found", zap.Error(err))
			c.JSON(http.StatusNotFound, gin.H{"error": "Posts not found"})
			return
		}
		pc.logger.Error("failed to get posts", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
	pc.logger.Info("posts fetched successfully")
}

// SearchPosts godoc
//
//	@Summary		Search posts
//	@Description	Search for posts by query string
//	@Tags			posts
//	@Produce		json
//	@Param			query	query		string					true	"Search query"
//	@Success		200		{object}	map[string]interface{}	"Posts searched successfully"
//	@Failure		400		{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		500		{object}	map[string]interface{}	"Internal server error"
//	@Securuty		BearerAuth
//	@Router			/post/search [get]
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
	pc.logger.Info("posts searched successfully", zap.String("query", query))
}

// GetPostWithPagination godoc
//
//	@Summary		Get posts with pagination
//	@Description	Retrieve paginated list of posts
//	@Tags			posts
//	@Produce		json
//	@Param			page	query		int						true	"Page number"
//	@Param			limit	query		int						true	"Items per page"
//	@Success		200		{object}	map[string]interface{}	"Posts fetched successfully"
//	@Failure		400		{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		500		{object}	map[string]interface{}	"Internal server error"
//	@Securuty		BearerAuth
//	@Router			/post/pagination [get]
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
	c.JSON(http.StatusOK, gin.H{"posts": posts, "totalPages": total, "page": pageInt})
	pc.logger.Info("posts fetched successfully with pagination", zap.Int("page", pageInt), zap.Int("limit", limitInt))
}
