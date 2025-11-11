package post_services

import (
	"flower-backend/models"

	"go.uber.org/zap"
)

// UpdatePostByID
func (s *PostService) UpdatePostByID(id uint, post models.Post) (*models.Post, error) {
	if err := s.db.Save(&post).Error; err != nil {
		s.logger.Error("failed to update post", zap.Error(err))
		return nil, err
	}
	s.logger.Info("post updated successfully", zap.Uint("id", id))
	return &post, nil
}

// UpdatePostByIDWithSelect
func (s *PostService) UpdatePostByIDWithSelect(id uint, post models.Post, selectFields []string) (*models.Post, error) {
	if err := s.db.Select(selectFields).Where("id = ?", id).Save(&post).Error; err != nil {
		s.logger.Error("failed to update post", zap.Error(err))
		return nil, err
	}
	s.logger.Info("post updated successfully", zap.Uint("id", id))
	return &post, nil
}
