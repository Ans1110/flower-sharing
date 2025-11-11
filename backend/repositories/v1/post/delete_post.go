package post_repository

import (
	"flower-backend/models"

	"go.uber.org/zap"
)

func (r *postRepository) DeleteByID(id uint) error {
	if err := r.db.Delete(&models.Post{}, id).Error; err != nil {
		r.logger.Error("failed to delete post by id", zap.Error(err))
		return err
	}
	return nil
}

