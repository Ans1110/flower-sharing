package user_services

import (
	"flower-backend/config"
	"flower-backend/log"
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct {
	db     *gorm.DB
	cfg    *config.Config
	user   models.User
	logger *zap.SugaredLogger
}

func NewUserService(db *gorm.DB, cfg *config.Config) *UserService {
	return &UserService{db: db, cfg: cfg, user: models.User{}, logger: log.InitLog().Sugar()}
}
