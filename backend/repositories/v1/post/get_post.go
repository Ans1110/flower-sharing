package post_repository

import (
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *postRepository) GetByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := r.db.Where("id = ?", id).First(&post).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to get post by id", zap.Error(err))
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) GetAllByUserID(userID uint) ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to get posts by user id", zap.Error(err))
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) GetAll() ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Find(&posts).Error; err != nil {
		r.logger.Error("failed to get all posts", zap.Error(err))
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) Search(query string) ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Where("title LIKE ? OR content LIKE ?", "%"+query+"%", "%"+query+"%").
		Order("created_at DESC").
		Find(&posts).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to search posts", zap.Error(err))
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) GetWithPagination(page, limit int) ([]models.Post, int64, error) {
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

	err := r.db.Table("posts").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error

	if err != nil {
		r.logger.Error("failed to get posts with pagination", zap.Error(err))
		return nil, 0, err
	}
	return posts, total, nil
}
