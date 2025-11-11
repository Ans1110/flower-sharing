package post_repository

import (
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *models.Post) error
	GetByID(id uint) (*models.Post, error)
	GetAllByUserID(userID uint) ([]models.Post, error)
	GetAll() ([]models.Post, error)
	Search(query string) ([]models.Post, error)
	GetWithPagination(page, limit int) ([]models.Post, int64, error)
	Update(post *models.Post) error
	UpdateByIDWithSelect(id uint, post *models.Post, selectFields []string) error
	DeleteByID(id uint) error
	Like(postID, userID uint) error
	Unlike(postID, userID uint) error
	CheckLikeExists(postID, userID uint) (bool, error)
	GetLikesCount(postID uint) (int64, error)
	GetUserLikedPosts(userID uint, page, limit int) ([]models.Post, int64, error)
}

type postRepository struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewPostRepository(db *gorm.DB, logger *zap.SugaredLogger) PostRepository {
	return &postRepository{
		db:     db,
		logger: logger,
	}
}
