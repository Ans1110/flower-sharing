package user_services

import (
	"errors"
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// FollowUser
func (s *UserService) FollowUser(followerID, followingID uint) error {
	exists, err := s.repo.CheckFollowExists(followerID, followingID)
	if err != nil {
		s.logger.Error("failed to check if user is followed", zap.Error(err))
		return err
	}
	if exists {
		return errors.New("user already followed")
	}

	if err := s.repo.Follow(followerID, followingID); err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.Error(err))
			return gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to follow user", zap.Error(err))
		return err
	}

	s.logger.Info("user followed successfully", zap.Uint("follower_id", followerID), zap.Uint("following_id", followingID))
	return nil
}

// UnfollowUser
func (s *UserService) UnfollowUser(followerID, followingID uint) error {
	if err := s.repo.Unfollow(followerID, followingID); err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.Error(err))
			return gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to unfollow user", zap.Error(err))
		return err
	}

	s.logger.Info("user unfollowed successfully", zap.Uint("follower_id", followerID), zap.Uint("following_id", followingID))
	return nil
}

// GetUserFollowers
func (s *UserService) GetUserFollowers(userID uint) ([]models.User, error) {
	followers, err := s.repo.GetFollowers(userID)
	if err != nil {
		s.logger.Error("failed to get user followers", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user followers fetched successfully", zap.Uint("user_id", userID))
	return followers, nil
}

// GetUserFollowing
func (s *UserService) GetUserFollowing(userID uint) ([]models.User, error) {
	following, err := s.repo.GetFollowing(userID)
	if err != nil {
		s.logger.Error("failed to get user following", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user following fetched successfully", zap.Uint("user_id", userID))
	return following, nil
}

// GetUserFollowersCount
func (s *UserService) GetUserFollowersCount(userID uint) (int64, error) {
	count, err := s.repo.GetFollowersCount(userID)
	if err != nil {
		s.logger.Error("failed to get user followers count", zap.Error(err))
		return 0, err
	}
	s.logger.Info("user followers count fetched successfully", zap.Uint("user_id", userID))
	return count, nil
}

// GetUserFollowingCount
func (s *UserService) GetUserFollowingCount(userID uint) (int64, error) {
	count, err := s.repo.GetFollowingCount(userID)
	if err != nil {
		s.logger.Error("failed to get user following count", zap.Error(err))
		return 0, err
	}
	s.logger.Info("user following count fetched successfully", zap.Uint("user_id", userID))
	return count, nil
}

// GetUserFollowingPosts
func (s *UserService) GetUserFollowingPosts(userID uint, page, limit int) ([]models.Post, int64, error) {
	posts, total, err := s.repo.GetFollowingPosts(userID, page, limit)
	if err != nil {
		s.logger.Error("failed to get user following posts", zap.Error(err))
		return nil, 0, err
	}
	s.logger.Info("user following posts fetched successfully", zap.Uint("user_id", userID))
	return posts, total, nil
}
