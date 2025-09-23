package services

import (
	"flower-backend/database"
	"flower-backend/models"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	return &user, err
}

func (s *UserService) GetUserByIDSelect(id uint, fields []string) (*models.User, error) {
	var user models.User
	err := database.DB.Select(fields).First(&user, id).Error
	return &user, err
}

func (s *UserService) GetUserUsername(id uint) (string, error) {
	user, err := s.GetUserByIDSelect(id, []string{"username"})
	if err != nil {
		return "", err
	}
	return user.Username, nil
}
