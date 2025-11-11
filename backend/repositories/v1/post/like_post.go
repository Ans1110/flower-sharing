package post_repository

import (
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *postRepository) Like(postID, userID uint) error {
	var post models.Post
	if err := r.db.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to find post", zap.Error(err))
		return err
	}

	if err := r.db.Model(&post).Association("Likes").Append(&models.User{ID: userID}); err != nil {
		r.logger.Error("failed to like post", zap.Error(err))
		return err
	}
	return nil
}

func (r *postRepository) Unlike(postID, userID uint) error {
	var post models.Post
	if err := r.db.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to find post", zap.Error(err))
		return err
	}

	if err := r.db.Model(&post).Association("Likes").Delete(&models.User{ID: userID}); err != nil {
		r.logger.Error("failed to unlike post", zap.Error(err))
		return err
	}
	return nil
}

func (r *postRepository) CheckLikeExists(postID, userID uint) (bool, error) {
	var count int64
	err := r.db.Table("post_likes").
		Where("post_id = ? AND user_id = ?", postID, userID).
		Count(&count).Error
	if err != nil {
		r.logger.Error("failed to check if post is liked", zap.Error(err))
		return false, err
	}
	return count > 0, nil
}

func (r *postRepository) GetLikesCount(postID uint) (int64, error) {
	var count int64
	if err := r.db.Table("post_likes").Where("post_id = ?", postID).Count(&count).Error; err != nil {
		r.logger.Error("failed to get post likes", zap.Error(err))
		return 0, err
	}
	return count, nil
}

func (r *postRepository) GetUserLikedPosts(userID uint, page, limit int) ([]models.Post, int64, error) {
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
		Joins("JOIN post_likes ON post_likes.post_id = posts.id").
		Where("post_likes.user_id = ?", userID).
		Preload("User").
		Order("posts.created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error

	if err != nil {
		r.logger.Error("failed to get user liked posts", zap.Error(err))
		return nil, 0, err
	}
	return posts, total, nil
}

