package public_user_controller

import (
	publicuserdto "flower-backend/dto/public"
	"flower-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// FollowUser
//
//	@Summary		Follow a user
//	@Description	Follow another user
//	@Tags			users
//	@Produce		json
//	@Param			follower_id		path		int	true	"Follower user ID"
//	@Param			following_id	path		int	true	"User to follow ID"
//	@Success		200				{object}	map[string]interface{}
//	@Failure		400				{object}	map[string]interface{}
//	@Failure		404				{object}	map[string]interface{}
//	@Failure		500				{object}	map[string]interface{}
//	@Securuty		BearerAuth
//	@Security		BearerAuth
//	@Router			/user/follow/{follower_id}/{following_id} [post]
func (uc *userController) FollowUser(c *gin.Context) {
	followerID := c.Param("follower_id")
	followingID := c.Param("following_id")
	followerIDUint, err := utils.ParseUint(followerID, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	followingIDUint, err := utils.ParseUint(followingID, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.svc.FollowUser(uint(followerIDUint), uint(followingIDUint)); err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.logger.Error("user not found", zap.String("follower_id", followerID), zap.String("following_id", followingID))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User followed successfully"})
	uc.logger.Info("user followed successfully", zap.String("follower_id", followerID), zap.String("following_id", followingID))
}

// UnfollowUser godoc
//
//	@Summary		Unfollow a user
//	@Description	Unfollow another user
//	@Tags			users
//	@Produce		json
//	@Param			follower_id		path		int	true	"Follower user ID"
//	@Param			following_id	path		int	true	"User to unfollow ID"
//	@Success		200				{object}	map[string]interface{}
//	@Failure		400				{object}	map[string]interface{}
func (uc *userController) UnfollowUser(c *gin.Context) {
	followerID := c.Param("follower_id")
	followingID := c.Param("following_id")
	followerIDUint, err := utils.ParseUint(followerID, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	followingIDUint, err := utils.ParseUint(followingID, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.svc.UnfollowUser(uint(followerIDUint), uint(followingIDUint)); err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.logger.Error("user not found", zap.String("follower_id", followerID), zap.String("following_id", followingID))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User unfollowed successfully"})
	uc.logger.Info("user unfollowed successfully", zap.String("follower_id", followerID), zap.String("following_id", followingID))
}

// GetUserFollowers godoc
//
//	@Summary		Get user followers
//	@Description	Get list of users following this user
//	@Tags			users
//	@Produce		json
//	@Param			user_id	path		int	true	"User ID"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		400		{object}	map[string]interface{}
//	@Failure		500		{object}	map[string]interface{}
//	@Securuty		BearerAuth
//	@Router			/user/followers/{user_id} [get]
func (uc *userController) GetUserFollowers(c *gin.Context) {
	userID := c.Param("user_id")
	userIDUint, err := utils.ParseUint(userID, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	followers, err := uc.svc.GetUserFollowers(uint(userIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"followers": publicuserdto.ToPublicUsers(followers)})
	uc.logger.Info("user followers fetched successfully", zap.String("user_id", userID))
}

// GetUserFollowing godoc
//
//	@Summary		Get user following
//	@Description	Get list of users followed by this user
//	@Tags			users
//	@Produce		json
//	@Param			user_id	path		int	true	"User ID"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		400		{object}	map[string]interface{}
//	@Failure		500		{object}	map[string]interface{}
//	@Securuty		BearerAuth
//	@Router			/user/following/{user_id} [get]
func (uc *userController) GetUserFollowing(c *gin.Context) {
	userID := c.Param("user_id")
	userIDUint, err := utils.ParseUint(userID, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	following, err := uc.svc.GetUserFollowing(uint(userIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"following": publicuserdto.ToPublicUsers(following)})
	uc.logger.Info("user following fetched successfully", zap.String("user_id", userID))
}

// GetUserFollowersCount godoc
//
//	@Summary		Get user followers count
//	@Description	Get the number of users following this user
//	@Tags			users
//	@Produce		json
//	@Param			user_id	path		int	true	"User ID"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		400		{object}	map[string]interface{}
//	@Failure		500		{object}	map[string]interface{}
//	@Securuty		BearerAuth
//	@Router			/user/followers-count/{user_id} [get]
func (uc *userController) GetUserFollowersCount(c *gin.Context) {
	userID := c.Param("user_id")
	userIDUint, err := utils.ParseUint(userID, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	count, err := uc.svc.GetUserFollowersCount(uint(userIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"followers_count": count})
	uc.logger.Info("user followers count fetched successfully", zap.String("user_id", userID))
}

// GetUserFollowingCount godoc
//
//	@Summary		Get user following count
//	@Description	Get the number of users followed by this user
//	@Tags			users
//	@Produce		json
//	@Param			user_id	path		int	true	"User ID"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		400		{object}	map[string]interface{}
//	@Failure		500		{object}	map[string]interface{}
//	@Securuty		BearerAuth
//	@Router			/user/following-count/{user_id} [get]
func (uc *userController) GetUserFollowingCount(c *gin.Context) {
	userID := c.Param("user_id")
	userIDUint, err := utils.ParseUint(userID, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	count, err := uc.svc.GetUserFollowingCount(uint(userIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"following_count": count})
	uc.logger.Info("user following count fetched successfully", zap.String("user_id", userID))
}

// GetUserFollowingPosts godoc
//
//	@Summary		Get user following posts
//	@Description	Get posts from users that this user follows
//	@Tags			users
//	@Produce		json
//	@Param			user_id	path		int	true	"User ID"
//	@Param			page	query		int	true	"Page number"
//	@Param			limit	query		int	true	"Items per page"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		400		{object}	map[string]interface{}
//	@Failure		500		{object}	map[string]interface{}
//	@Securuty		BearerAuth
//	@Router			/user/following-posts/{user_id} [get]
func (uc *userController) GetUserFollowingPosts(c *gin.Context) {
	userID := c.Param("user_id")
	userIDUint, err := utils.ParseUint(userID, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page := c.Query("page")
	limit := c.Query("limit")
	pageUint, err := utils.ParseUint(page, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limitUint, err := utils.ParseUint(limit, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	posts, total, err := uc.svc.GetUserFollowingPosts(uint(userIDUint), int(pageUint), int(limitUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": publicuserdto.ToPublicPosts(posts), "total": total})
	uc.logger.Info("user following posts fetched successfully", zap.String("user_id", userID))
}
