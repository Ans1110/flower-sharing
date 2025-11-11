package post_services

import (
	"flower-backend/config"
	"flower-backend/log"
	post_repository "flower-backend/repositories/v1/post"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PostService struct {
	repo   post_repository.PostRepository
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewPostService(db *gorm.DB, cfg *config.Config) *PostService {
	logger := log.InitLog().Sugar()
	repo := post_repository.NewPostRepository(db, logger)
	return &PostService{repo: repo, cfg: cfg, logger: logger}
}
