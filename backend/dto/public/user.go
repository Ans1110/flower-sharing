package public_dto

import (
	"strings"

	"flower-backend/models"
)

type PublicUserDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func ToPublicUser(user *models.User) PublicUserDTO {
	if user == nil {
		return PublicUserDTO{}
	}

	return PublicUserDTO{
		ID:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
	}
}

func ToPublicUsers(users []models.User) []PublicUserDTO {
	result := make([]PublicUserDTO, 0, len(users))
	for i := range users {
		result = append(result, ToPublicUser(&users[i]))
	}
	return result
}

func EnsurePublicUserSelectFields(fields []string) []string {
	required := map[string]struct{}{
		"id":       {},
		"username": {},
		"avatar":   {},
	}

	normalized := make(map[string]struct{}, len(fields))
	cleaned := make([]string, 0, len(fields))
	for _, field := range fields {
		trimmed := strings.TrimSpace(field)
		if trimmed == "" {
			continue
		}
		if _, exists := normalized[trimmed]; exists {
			continue
		}
		normalized[trimmed] = struct{}{}
		cleaned = append(cleaned, trimmed)
	}

	for field := range required {
		if _, ok := normalized[field]; !ok {
			cleaned = append(cleaned, field)
		}
	}

	return cleaned
}
