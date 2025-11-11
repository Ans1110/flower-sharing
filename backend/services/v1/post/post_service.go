package post_services

import (
	"flower-backend/config"
	"flower-backend/log"
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PostService struct {
	db     *gorm.DB
	cfg    *config.Config
	post   models.Post
	logger *zap.SugaredLogger
}

func NewPostService(db *gorm.DB, cfg *config.Config) *PostService {
	return &PostService{db: db, cfg: cfg, post: models.Post{}, logger: log.InitLog().Sugar()}
}
