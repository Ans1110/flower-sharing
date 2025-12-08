package post_repository

import (
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *postRepository) Like(postID, userID uint) error {
	// Check if post exists
	var postCount int64
	if err := r.db.Model(&models.Post{}).Where("id = ?", postID).Count(&postCount).Error; err != nil {
		r.logger.Error("failed to check if post exists", zap.Error(err))
		return err
	}
	if postCount == 0 {
		return gorm.ErrRecordNotFound
	}

	// Check if user exists
	var userCount int64
	if err := r.db.Model(&models.User{}).Where("id = ?", userID).Count(&userCount).Error; err != nil {
		r.logger.Error("failed to check if user exists", zap.Error(err))
		return err
	}
	if userCount == 0 {
		r.logger.Error("user not found", zap.Uint("user_id", userID))
		return gorm.ErrRecordNotFound
	}

	// Insert directly into the join table using raw SQL
	// This is more efficient than loading full objects
	// Using INSERT IGNORE for MySQL to handle potential race conditions
	result := r.db.Exec(
		"INSERT IGNORE INTO post_likes (post_id, user_id) VALUES (?, ?)",
		postID, userID,
	)
	if result.Error != nil {
		r.logger.Error("failed to like post",
			zap.Uint("post_id", postID),
			zap.Uint("user_id", userID),
			zap.Error(result.Error))
		return result.Error
	}

	// INSERT IGNORE returns 0 rows affected if duplicate exists (race condition)
	// This is fine - the service layer already checks for duplicates before calling this
	// If we get here and RowsAffected is 0, it means another request inserted it first
	// which is acceptable behavior (idempotent operation)
	return nil
}

func (r *postRepository) Unlike(postID, userID uint) error {
	// Delete directly from the join table using raw SQL
	// This is more efficient than loading full objects
	result := r.db.Exec(
		"DELETE FROM post_likes WHERE post_id = ? AND user_id = ?",
		postID, userID,
	)
	if result.Error != nil {
		r.logger.Error("failed to unlike post", zap.Error(result.Error))
		return result.Error
	}

	// Check if any rows were affected (post or like might not exist)
	if result.RowsAffected == 0 {
		// Check if post exists
		var postCount int64
		if err := r.db.Model(&models.Post{}).Where("id = ?", postID).Count(&postCount).Error; err != nil {
			r.logger.Error("failed to check if post exists", zap.Error(err))
			return err
		}
		if postCount == 0 {
			return gorm.ErrRecordNotFound
		}
		// Post exists but like doesn't - this is fine, just return success
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
		Preload("Likes").
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
