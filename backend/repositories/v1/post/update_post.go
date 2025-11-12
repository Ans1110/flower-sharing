package post_repository

import (
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *postRepository) UpdateByIDWithSelect(postId uint, updates map[string]any, selectFields []string) (*models.Post, error) {
	var post models.Post

	if err := r.db.First(&post, postId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to find post", zap.Error(err))
		return nil, err
	}

	filtered := make(map[string]any)
	for _, field := range selectFields {
		if val, ok := updates[field]; ok {
			filtered[field] = val
		}
	}

	if len(filtered) == 0 {
		return &post, nil
	}

	if err := r.db.Model(&post).Updates(filtered).Error; err != nil {
		r.logger.Error("failed to update post", zap.Error(err))
		return nil, err
	}

	if err := r.db.First(&post, postId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to find post", zap.Error(err))
		return nil, err
	}

	r.logger.Info("post updated successfully", zap.Uint("id", postId))
	return &post, nil
}

func (r *postRepository) Update(post *models.Post) error {
	if err := r.db.Save(post).Error; err != nil {
		r.logger.Error("failed to update post", zap.Error(err))
		return err
	}
	return nil
}
