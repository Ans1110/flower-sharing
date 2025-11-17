package user_repository

import (
	"flower-backend/config"
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	CreateToken(token *models.Token) error
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetByIDWithSelect(id uint, selectFields []string) (*models.User, error)
	GetAll() ([]models.User, error)
	Update(user *models.User) error
	UpdateByIDWithSelect(id uint, updates map[string]any, selectFields []string) (*models.User, error)
	DeleteByID(id uint) error
	Follow(followerID, followingID uint) error
	Unfollow(followerID, followingID uint) error
	CheckFollowExists(followerID, followingID uint) (bool, error)
	GetFollowers(userID uint) ([]models.User, error)
	GetFollowing(userID uint) ([]models.User, error)
	GetFollowersCount(userID uint) (int64, error)
	GetFollowingCount(userID uint) (int64, error)
	GetFollowingPosts(userID uint, page, limit int) ([]models.Post, int64, error)
}

type userRepository struct {
	db     *gorm.DB
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewUserRepository(db *gorm.DB, cfg *config.Config, logger *zap.SugaredLogger) UserRepository {
	return &userRepository{
		db:     db,
		cfg:    cfg,
		logger: logger,
	}
}
