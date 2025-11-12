package post_services

import (
	"flower-backend/config"
	"flower-backend/models"
	post_repository "flower-backend/repositories/v1/post"
	"mime/multipart"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PostService interface {
	CreatePost(post models.Post) (*models.Post, error)
	UploadImage(buffer []byte, postID uint) (string, error)
	GetPostByID(id uint) (*models.Post, error)
	GetPostAllByUserID(userID uint) ([]models.Post, error)
	GetPostAll() ([]models.Post, error)
	SearchPosts(query string) ([]models.Post, error)
	GetPostWithPagination(page, limit int) ([]models.Post, int64, error)
	CheckPostOwnership(postID, userID uint) (bool, error)
	UpdatePostByID(postId uint, userId uint, imageFile *multipart.FileHeader, updates map[string]any, selectFields []string) (*models.Post, error)
	DeletePostByID(postID, userID uint) error
	LikePost(postID, userID uint) error
	DislikePost(postID, userID uint) error
	GetPostLikes(postID uint) (int64, error)
	GetUserLikedPosts(userID uint, page, limit int) ([]models.Post, int64, error)
}

type postService struct {
	repo   post_repository.PostRepository
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewPostService(db *gorm.DB, cfg *config.Config, logger *zap.SugaredLogger) PostService {
	repo := post_repository.NewPostRepository(db, cfg, logger)
	return &postService{repo: repo, cfg: cfg, logger: logger}
}
