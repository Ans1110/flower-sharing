package post_repository

import (
	"flower-backend/models"

	"go.uber.org/zap"
)

func (r *postRepository) Create(post *models.Post) error {
	if err := r.db.Create(post).Error; err != nil {
		r.logger.Error("failed to create post", zap.Error(err))
		return err
	}
	return nil
}
