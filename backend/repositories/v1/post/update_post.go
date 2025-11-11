package post_repository

import (
	"flower-backend/models"

	"go.uber.org/zap"
)

func (r *postRepository) Update(post *models.Post) error {
	if err := r.db.Save(post).Error; err != nil {
		r.logger.Error("failed to update post", zap.Error(err))
		return err
	}
	return nil
}

func (r *postRepository) UpdateByIDWithSelect(id uint, post *models.Post, selectFields []string) error {
	if err := r.db.Select(selectFields).Where("id = ?", id).Save(post).Error; err != nil {
		r.logger.Error("failed to update post by id with select", zap.Error(err))
		return err
	}
	return nil
}

