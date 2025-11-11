package user_services

import (
	"errors"
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// FollowUser
func (s *UserService) FollowUser(followerID, followingID uint) error {
	var follower, following models.User
	if err := s.db.First(&follower, followerID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("follower not found", zap.Uint("id", followerID))
			return gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to follow user", zap.Error(err))
		return err
	}
	if err := s.db.First(&following, followingID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("following not found", zap.Uint("id", followingID))
			return gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to follow user", zap.Error(err))
		return err
	}

	var count int64
	err := s.db.Table("user_follows").
		Where("follower_id = ? AND following_id = ?", followerID, followingID).
		Count(&count).Error
	if err != nil {
		s.logger.Error("failed to check if user is followed", zap.Error(err))
		return err
	}
	if count > 0 {
		return errors.New("user already followed")
	}

	if err := s.db.Model(&follower).Association("Following").Append(&models.User{ID: followingID}); err != nil {
		s.logger.Error("failed to follow user", zap.Error(err))
		return err
	}

	if err := s.db.Model(&following).Association("Followers").Append(&models.User{ID: followerID}); err != nil {
		s.logger.Error("failed to follow user", zap.Error(err))
		return err
	}

	s.logger.Info("user followed successfully", zap.Uint("follower_id", followerID), zap.Uint("following_id", followingID))
	return nil
}

// UnfollowUser
func (s *UserService) UnfollowUser(followerID, followingID uint) error {
	var follower, following models.User
	if err := s.db.First(&follower, followerID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("follower not found", zap.Uint("id", followerID))
			return gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to unfollow user", zap.Error(err))
		return err
	}
	if err := s.db.First(&following, followingID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("following not found", zap.Uint("id", followingID))
			return gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to unfollow user", zap.Error(err))
		return err
	}

	if err := s.db.Model(&follower).Association("Following").Delete(&models.User{ID: followingID}); err != nil {
		s.logger.Error("failed to unfollow user", zap.Error(err))
		return err
	}

	if err := s.db.Model(&following).Association("Followers").Delete(&models.User{ID: followerID}); err != nil {
		s.logger.Error("failed to unfollow user", zap.Error(err))
		return err
	}

	s.logger.Info("user unfollowed successfully", zap.Uint("follower_id", followerID), zap.Uint("following_id", followingID))
	return nil
}

// GetUserFollowers
func (s *UserService) GetUserFollowers(userID uint) ([]models.User, error) {
	var followers []models.User
	if err := s.db.Model(&models.User{}).Where("following_id = ?", userID).Find(&followers).Error; err != nil {
		s.logger.Error("failed to get user followers", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user followers fetched successfully", zap.Uint("user_id", userID))
	return followers, nil
}

// GetUserFollowing
func (s *UserService) GetUserFollowing(userID uint) ([]models.User, error) {
	var following []models.User
	if err := s.db.Model(&models.User{}).Where("follower_id = ?", userID).Find(&following).Error; err != nil {
		s.logger.Error("failed to get user following", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user following fetched successfully", zap.Uint("user_id", userID))
	return following, nil
}

// GetUserFollowersCount
func (s *UserService) GetUserFollowersCount(userID uint) (int64, error) {
	var count int64
	if err := s.db.Table("user_follows").
		Where("following_id = ?", userID).
		Count(&count).Error; err != nil {
		s.logger.Error("failed to get user followers count", zap.Error(err))
		return 0, err
	}
	s.logger.Info("user followers count fetched successfully", zap.Uint("user_id", userID))
	return count, nil
}

// GetUserFollowingCount
func (s *UserService) GetUserFollowingCount(userID uint) (int64, error) {
	var count int64
	if err := s.db.Table("user_follows").
		Where("follower_id = ?", userID).
		Count(&count).Error; err != nil {
		s.logger.Error("failed to get user following count", zap.Error(err))
		return 0, err
	}
	s.logger.Info("user following count fetched successfully", zap.Uint("user_id", userID))
	return count, nil
}

// GetUserFollowingPosts
func (s *UserService) GetUserFollowingPosts(userID uint) ([]models.Post, error) {
	var posts []models.Post
	if err := s.db.Model(&models.Post{}).Where("user_id IN (?)", userID).Find(&posts).Error; err != nil {
		s.logger.Error("failed to get user following posts", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user following posts fetched successfully", zap.Uint("user_id", userID))
	return posts, nil
}
