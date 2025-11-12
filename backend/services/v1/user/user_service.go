package user_services

import (
	"flower-backend/config"
	"flower-backend/models"
	user_repository "flower-backend/repositories/v1/user"
	"mime/multipart"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user models.User) (*models.User, error)
	RegisterUser(username, email, password string, avatarFile *multipart.FileHeader) (*models.User, error)
	UploadAvatar(buffer []byte, userID uint) (string, error)
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByIDWithSelect(id uint, selectFields []string) (*models.User, error)
	GetUserAll() ([]models.User, error)
	UpdateUserByIDWithSelect(id uint, updates map[string]any, imageFile *multipart.FileHeader, selectFields []string) (*models.User, error)
	DeleteUserByID(id uint) error
	FollowUser(followerID, followingID uint) error
	UnfollowUser(followerID, followingID uint) error
	GetUserFollowers(userID uint) ([]models.User, error)
	GetUserFollowing(userID uint) ([]models.User, error)
	GetUserFollowersCount(userID uint) (int64, error)
	GetUserFollowingCount(userID uint) (int64, error)
	GetUserFollowingPosts(userID uint, page, limit int) ([]models.Post, int64, error)
}

type userService struct {
	repo   user_repository.UserRepository
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewUserService(db *gorm.DB, cfg *config.Config, logger *zap.SugaredLogger) UserService {
	repo := user_repository.NewUserRepository(db, cfg, logger)
	return &userService{repo: repo, cfg: cfg, logger: logger}
}
