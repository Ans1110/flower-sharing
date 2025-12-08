package user_repository

import (
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *userRepository) Follow(followerID, followingID uint) error {
	var follower models.User
	if err := r.db.First(&follower, followerID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to find follower", zap.Error(err))
		return err
	}

	if err := r.db.Model(&follower).Association("Following").Append(&models.User{ID: followingID}); err != nil {
		r.logger.Error("failed to follow user", zap.Error(err))
		return err
	}

	return nil
}

func (r *userRepository) Unfollow(followerID, followingID uint) error {
	var follower models.User
	if err := r.db.First(&follower, followerID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to find follower", zap.Error(err))
		return err
	}

	if err := r.db.Model(&follower).Association("Following").Delete(&models.User{ID: followingID}); err != nil {
		r.logger.Error("failed to unfollow user", zap.Error(err))
		return err
	}

	return nil
}

func (r *userRepository) CheckFollowExists(followerID, followingID uint) (bool, error) {
	var count int64
	err := r.db.Table("user_follows").
		Where("follower_id = ? AND following_id = ?", followerID, followingID).
		Count(&count).Error
	if err != nil {
		r.logger.Error("failed to check if user is followed", zap.Error(err))
		return false, err
	}
	return count > 0, nil
}

func (r *userRepository) GetFollowers(userID uint) ([]models.User, error) {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Warn("user not found when getting followers", zap.Uint("user_id", userID))
			return []models.User{}, nil
		}
		r.logger.Error("failed to find user", zap.Error(err))
		return nil, err
	}

	var followers []models.User
	if err := r.db.Model(&user).Association("Followers").Find(&followers); err != nil {
		r.logger.Error("failed to get user followers", zap.Error(err))
		return nil, err
	}
	return followers, nil
}

func (r *userRepository) GetFollowing(userID uint) ([]models.User, error) {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Warn("user not found when getting following", zap.Uint("user_id", userID))
			return []models.User{}, nil
		}
		r.logger.Error("failed to find user", zap.Error(err))
		return nil, err
	}

	var following []models.User
	if err := r.db.Model(&user).Association("Following").Find(&following); err != nil {
		r.logger.Error("failed to get user following", zap.Error(err))
		return nil, err
	}
	return following, nil
}

func (r *userRepository) GetFollowersCount(userID uint) (int64, error) {
	var count int64
	if err := r.db.Table("user_follows").
		Where("following_id = ?", userID).
		Count(&count).Error; err != nil {
		r.logger.Error("failed to get user followers count", zap.Error(err))
		return 0, err
	}
	return count, nil
}

func (r *userRepository) GetFollowingCount(userID uint) (int64, error) {
	var count int64
	if err := r.db.Table("user_follows").
		Where("follower_id = ?", userID).
		Count(&count).Error; err != nil {
		r.logger.Error("failed to get user following count", zap.Error(err))
		return 0, err
	}
	return count, nil
}

func (r *userRepository) GetFollowingPosts(userID uint, page, limit int) ([]models.Post, int64, error) {
	var posts []models.Post
	var total int64

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 6
	}

	offset := (page - 1) * limit
	if err := r.db.Table("posts").Count(&total).Error; err != nil {
		r.logger.Error("failed to get total posts", zap.Error(err))
		return nil, 0, err
	}

	err := r.db.
		Joins("JOIN user_follows ON user_follows.following_id = posts.user_id").
		Where("user_follows.follower_id = ?", userID).
		Preload("User").
		Order("posts.created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error
	if err != nil {
		r.logger.Error("failed to get user following posts", zap.Error(err))
		return nil, 0, err
	}
	return posts, total, nil
}
