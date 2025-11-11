package user_repository

import (
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetByIDWithSelect(id uint, selectFields []string) (*models.User, error)
	GetByEmailWithSelect(email string, selectFields []string) (*models.User, error)
	GetByUsernameWithSelect(username string, selectFields []string) (*models.User, error)
	GetAll() ([]models.User, error)
	Update(user *models.User) error
	UpdateByID(id uint, user *models.User) error
	UpdateByIDWithSelect(id uint, user *models.User, selectFields []string) error
	DeleteByID(id uint) error
	DeleteByEmail(email string) error
	DeleteByUsername(username string) error
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
	logger *zap.SugaredLogger
}

func NewUserRepository(db *gorm.DB, logger *zap.SugaredLogger) UserRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}
