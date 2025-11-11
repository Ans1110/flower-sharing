package post_services

import (
	"errors"
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GetPostByID
func (s *PostService) GetPostByID(id uint) (*models.Post, error) {
	post, err := s.repo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("post not found", zap.Uint("id", id))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get post by id", zap.Error(err))
		return nil, err
	}
	s.logger.Info("post fetched successfully", zap.Uint("id", id))
	return post, nil
}

// GetPostByUserID
func (s *PostService) GetPostAllByUserID(userID uint) ([]models.Post, error) {
	posts, err := s.repo.GetAllByUserID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("posts not found", zap.Uint("user_id", userID))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get posts by user id", zap.Error(err))
		return nil, err
	}
	s.logger.Info("posts fetched successfully", zap.Uint("user_id", userID))
	return posts, nil
}

// GetPostAll
func (s *PostService) GetPostAll() ([]models.Post, error) {
	posts, err := s.repo.GetAll()
	if err != nil {
		s.logger.Error("failed to get all posts", zap.Error(err))
		return nil, err
	}
	s.logger.Info("all posts fetched successfully")
	return posts, nil
}

// SearchPosts
func (s *PostService) SearchPosts(query string) ([]models.Post, error) {
	posts, err := s.repo.Search(query)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("posts not found", zap.String("query", query))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to search posts", zap.Error(err))
		return nil, err
	}
	s.logger.Info("posts searched successfully", zap.String("query", query))
	return posts, nil
}

// CheckPostOwnership
func (s *PostService) CheckPostOwnership(postID, userID uint) (bool, error) {
	post, err := s.repo.GetByID(postID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("post not found", zap.Uint("id", postID))
			return false, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to check post ownership", zap.Error(err))
		return false, err
	}

	if post.User.Role == "admin" {
		s.logger.Info("post owned by admin", zap.Uint("id", postID))
		return true, nil
	}

	if post.UserID != userID {
		s.logger.Error("post not owned by user", zap.Uint("id", postID))
		return false, errors.New("post not owned by user")
	}
	s.logger.Info("post ownership checked successfully", zap.Uint("id", postID))
	return true, nil
}

// GetPostWithPagination
func (s *PostService) GetPostWithPagination(page, limit int) ([]models.Post, int64, error) {
	posts, total, err := s.repo.GetWithPagination(page, limit)
	if err != nil {
		s.logger.Error("failed to get posts with pagination", zap.Error(err))
		return nil, 0, err
	}
	s.logger.Info("posts fetched successfully with pagination", zap.Int("page", page), zap.Int("limit", limit))
	return posts, total, nil
}
