package services

import (
	"flower-backend/config"
	"flower-backend/log"
	"flower-backend/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct {
	db     *gorm.DB
	cfg    *config.Config
	user   models.User
	logger *zap.SugaredLogger
}

func NewUserService(db *gorm.DB, cfg *config.Config) *UserService {
	return &UserService{db: db, cfg: cfg, user: models.User{}, logger: log.InitLog().Sugar()}
}

// CreateUser
func (s *UserService) CreateUser(user models.User) (*models.User, error) {
	if err := s.db.Create(&user).Error; err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user created successfully", zap.String("username", user.Username))
	return &user, nil
}

// GetUserByID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	if err := s.db.Where("id = ?", id).First(&s.user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.Uint("id", id))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by id", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.Uint("id", id))
	return &s.user, nil
}

// GetUserByEmail
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	if err := s.db.Where("email = ?", email).First(&s.user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.String("email", email))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by email", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.String("email", email))
	return &s.user, nil
}

// GetUserByUsername
func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	if err := s.db.Where("username = ?", username).First(&s.user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.String("username", username))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by username", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.String("username", username))
	return &s.user, nil
}

// GetUserByIDWithSelect
func (s *UserService) GetUserByIDWithSelect(id uint, selectFields []string) (*models.User, error) {
	if err := s.db.Select(selectFields).Where("id = ?", id).First(&s.user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.Uint("id", id))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by id", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.Uint("id", id))
	return &s.user, nil
}

// GetUserByEmailWithSelect
func (s *UserService) GetUserByEmailWithSelect(email string, selectFields []string) (*models.User, error) {
	if err := s.db.Select(selectFields).Where("email = ?", email).First(&s.user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.String("email", email))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by email", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.String("email", email))
	return &s.user, nil
}

// GetUserByUsernameWithSelect
func (s *UserService) GetUserByUsernameWithSelect(username string, selectFields []string) (*models.User, error) {
	if err := s.db.Select(selectFields).Where("username = ?", username).First(&s.user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Error("user not found", zap.String("username", username))
			return nil, gorm.ErrRecordNotFound
		}
		s.logger.Error("failed to get user by username", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user fetched successfully", zap.String("username", username))
	return &s.user, nil
}

// GetUserAll
func (s *UserService) GetUserAll() ([]models.User, error) {
	var users []models.User
	if err := s.db.Find(&users).Error; err != nil {
		s.logger.Error("failed to get all users", zap.Error(err))
		return nil, err
	}
	s.logger.Info("all users fetched successfully")
	return users, nil
}

// UpdateUser
func (s *UserService) UpdateUser(user models.User) (*models.User, error) {
	if err := s.db.Save(&user).Error; err != nil {
		s.logger.Error("failed to update user", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user updated successfully", zap.String("username", user.Username))
	return &user, nil
}

// UpdateUserWithSelect
func (s *UserService) UpdateUserWithSelect(user models.User, selectFields []string) (*models.User, error) {
	if err := s.db.Select(selectFields).Save(&user).Error; err != nil {
		s.logger.Error("failed to update user", zap.Error(err))
		return nil, err
	}
	s.logger.Info("user updated successfully", zap.String("username", user.Username))
	return &user, nil
}

// DeleteUserByID
func (s *UserService) DeleteUserByID(id uint) error {
	if err := s.db.Delete(&models.User{}, id).Error; err != nil {
		s.logger.Error("failed to delete user", zap.Error(err))
		return err
	}
	s.logger.Info("user deleted successfully", zap.Uint("id", id))
	return nil
}

// DeleteUserByEmail
func (s *UserService) DeleteUserByEmail(email string) error {
	if err := s.db.Delete(&models.User{}, "email = ?", email).Error; err != nil {
		s.logger.Error("failed to delete user", zap.Error(err))
		return err
	}
	s.logger.Info("user deleted successfully", zap.String("email", email))
	return nil
}

// DeleteUserByUsername
func (s *UserService) DeleteUserByUsername(username string) error {
	if err := s.db.Delete(&models.User{}, "username = ?", username).Error; err != nil {
		s.logger.Error("failed to delete user", zap.Error(err))
		return err
	}
	s.logger.Info("user deleted successfully", zap.String("username", username))
	return nil
}
