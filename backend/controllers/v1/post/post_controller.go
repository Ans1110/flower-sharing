package post_controller

import (
	"flower-backend/config"
	post_services "flower-backend/services/v1/post"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PostController interface {
	CreatePost(c *gin.Context)
	GetPostByID(c *gin.Context)
	GetPostAllByUserID(c *gin.Context)
	GetPostAll(c *gin.Context)
	SearchPosts(c *gin.Context)
	GetPostWithPagination(c *gin.Context)
	UpdatePostByIDWithSelect(c *gin.Context)
	DeletePostByID(c *gin.Context)
	LikePost(c *gin.Context)
	DislikePost(c *gin.Context)
	GetPostLikes(c *gin.Context)
	GetUserLikedPosts(c *gin.Context)
}

type postController struct {
	svc    post_services.PostService
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func NewPostController(db *gorm.DB, cfg *config.Config, logger *zap.SugaredLogger) PostController {
	svc := post_services.NewPostService(db, cfg, logger)
	return &postController{svc: svc, logger: logger, cfg: cfg}
}
