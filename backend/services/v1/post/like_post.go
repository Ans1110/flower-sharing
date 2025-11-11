package post_services

import (
	"errors"
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// LikePost
func (s *PostService) LikePost(postID, userID uint) error {
	exists, err := s.repo.CheckLikeExists(postID, userID)
	if err != nil {
		s.logger.Error("failed to check if post is liked", zap.Error(err))
		return err
	}

	if exists {
		return errors.New("post already liked")
	}

	if err := s.repo.Like(postID, userID); err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("post not found", zap.Uint("id", postID))
			return gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to like post", zap.Error(err))
		return err
	}

	s.logger.Info("post liked successfully", zap.Uint("post_id", postID), zap.Uint("user_id", userID))
	return nil
}

// UnlikePost
func (s *PostService) UnlikePost(postID, userID uint) error {
	if err := s.repo.Unlike(postID, userID); err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("post not found", zap.Uint("id", postID))
			return gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to unlike post", zap.Error(err))
		return err
	}

	s.logger.Info("post unliked successfully", zap.Uint("post_id", postID), zap.Uint("user_id", userID))
	return nil
}

// GetPostLikes
func (s *PostService) GetPostLikes(postID uint) (int64, error) {
	count, err := s.repo.GetLikesCount(postID)
	if err != nil {
		s.logger.Error("failed to get post likes", zap.Error(err))
		return 0, err
	}
	s.logger.Info("post likes fetched successfully", zap.Uint("post_id", postID))
	return count, nil
}

// GetUserLikedPosts
func (s *PostService) GetUserLikedPosts(userID uint, page, limit int) ([]models.Post, int64, error) {
	posts, total, err := s.repo.GetUserLikedPosts(userID, page, limit)
	if err != nil {
		s.logger.Error("failed to get user liked posts", zap.Error(err))
		return nil, 0, err
	}
	s.logger.Info("user liked posts fetched successfully", zap.Uint("user_id", userID))
	return posts, total, nil
}
