package user_services

import (
	"flower-backend/config"
	"flower-backend/log"
	user_repository "flower-backend/repositories/v1/user"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct {
	repo   user_repository.UserRepository
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewUserService(db *gorm.DB, cfg *config.Config) *UserService {
	logger := log.InitLog().Sugar()
	repo := user_repository.NewUserRepository(db, logger)
	return &UserService{repo: repo, cfg: cfg, logger: logger}
}
