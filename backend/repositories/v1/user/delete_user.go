package user_repository

import (
	"flower-backend/libs"
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *userRepository) DeleteByID(id uint) error {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		r.logger.Error("failed to find user", zap.Error(err))
		return err
	}

	if user.Avatar != "" {
		cld, _ := libs.NewCloudinary(r.cfg)
		publicId := libs.ExtractPublicId(user.Avatar)
		if err := libs.DeleteFromCloudinary(cld, publicId); err != nil {
			r.logger.Error("failed to delete avatar from cloudinary", zap.Error(err))
			return err
		}
	}

	if err := r.db.Where("user_id = ?", id).Delete(&models.Token{}).Error; err != nil {
		r.logger.Error("failed to delete user tokens", zap.Error(err))
		return err
	}

	if err := r.db.Delete(&user).Error; err != nil {
		r.logger.Error("failed to delete user by id", zap.Error(err))
		return err
	}
	return nil
}
