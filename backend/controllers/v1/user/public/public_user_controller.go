package public_user_controller

import (
	"flower-backend/config"
	user_services "flower-backend/services/v1/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserController interface {
	GetUserByID(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	GetUserByUsername(c *gin.Context)
	GetUserAll(c *gin.Context)
	GetUserByIDWithSelect(c *gin.Context)
	UpdateUserByIDWithSelect(c *gin.Context)
	FollowUser(c *gin.Context)
	UnfollowUser(c *gin.Context)
	GetUserFollowers(c *gin.Context)
	GetUserFollowing(c *gin.Context)
	GetUserFollowersCount(c *gin.Context)
	GetUserFollowingCount(c *gin.Context)
	GetUserFollowingPosts(c *gin.Context)
}

type userController struct {
	svc    user_services.UserService
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewUserController(db *gorm.DB, cfg *config.Config, logger *zap.SugaredLogger) UserController {
	svc := user_services.NewUserService(db, cfg, logger)
	return &userController{svc: svc, logger: logger, cfg: cfg}
}
