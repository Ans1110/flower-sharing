package post_repository

import (
	"flower-backend/libs"
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *postRepository) DeleteByID(postID, userID uint) error {
	var post models.Post
	if err := r.db.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to find post", zap.Error(err))
		return err
	}

	if post.ImageURL != "" {
		cld, _ := libs.NewCloudinary(r.cfg)
		publicId := libs.ExtractPublicId(post.ImageURL)
		if err := libs.DeleteFromCloudinary(cld, publicId); err != nil {
			r.logger.Error("failed to delete image from cloudinary", zap.Error(err))
			return err
		}
	}

	if err := r.db.Delete(&post).Error; err != nil {
		r.logger.Error("failed to delete post", zap.Error(err))
		return err
	}
	return nil
}
