package public_user_controller

import (
	publicuserdto "flower-backend/dto/public"
	"flower-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// POST /api/v1/user/follow/:follower_id/:following_id
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

// POST /api/v1/user/unfollow/:follower_id/:following_id
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

// GET /api/v1/user/followers/:user_id
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

// GET /api/v1/user/following/:user_id
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

// GET /api/v1/user/followers-count/:user_id
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

// GET /api/v1/user/following-count/:user_id
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

// GET /api/v1/user/following-posts/:user_id
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
