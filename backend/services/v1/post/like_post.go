package post_services

import (
	"errors"
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// LikePost
func (s *PostService) LikePost(postID, userID uint) error {
	var post models.Post
	if err := s.db.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("post not found", zap.Uint("id", postID))
			return gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to like post", zap.Error(err))
		return err
	}

	var count int64
	err := s.db.Table("post_likes").
		Where("post_id = ? AND user_id = ?", postID, userID).
		Count(&count).Error

	if err != nil {
		s.logger.Error("failed to check if post is liked", zap.Error(err))
		return err
	}

	if count > 0 {
		return errors.New("post already liked")
	}

	if err := s.db.Model(&post).Association("Likes").Append(&models.User{ID: userID}); err != nil {
		s.logger.Error("failed to like post", zap.Error(err))
		return err
	}

	s.logger.Info("post liked successfully", zap.Uint("post_id", postID), zap.Uint("user_id", userID))
	return nil
}

// UnlikePost
func (s *PostService) UnlikePost(postID, userID uint) error {
	var post models.Post
	if err := s.db.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("post not found", zap.Uint("id", postID))
			return gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to unlike post", zap.Error(err))
		return err
	}

	if err := s.db.Model(&post).Association("Likes").Delete(&models.User{ID: userID}); err != nil {
		s.logger.Error("failed to unlike post", zap.Error(err))
		return err
	}

	s.logger.Info("post unliked successfully", zap.Uint("post_id", postID), zap.Uint("user_id", userID))
	return nil
}

// GetPostLikes
func (s *PostService) GetPostLikes(postID uint) (int64, error) {
	var count int64
	if err := s.db.Table("post_likes").Where("post_id = ?", postID).Count(&count).Error; err != nil {
		s.logger.Error("failed to get post likes", zap.Error(err))
		return 0, err
	}
	s.logger.Info("post likes fetched successfully", zap.Uint("post_id", postID))
	return count, nil
}

// GetUserLikedPosts
func (s *PostService) GetUserLikedPosts(userID uint, page, limit int) ([]models.Post, int64, error) {
	var posts []models.Post
	var total int64

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 6
	}

	offset := (page - 1) * limit
	if err := s.db.Table("posts").Count(&total).Error; err != nil {
		s.logger.Error("failed to get total posts", zap.Error(err))
		return nil, 0, err
	}

	err := s.db.
		Joins("JOIN post_likes ON post_likes.post_id = posts.id").
		Where("post_likes.user_id = ?", userID).
		Preload("User").
		Order("posts.created_at DESC").
		Offset(offset).
		Limit(limit).
		Scan(&posts).Error

	if err != nil {
		s.logger.Error("failed to get user liked posts", zap.Error(err))
		return nil, 0, err
	}
	s.logger.Info("user liked posts fetched successfully", zap.Uint("user_id", userID))
	return posts, total, nil
}
