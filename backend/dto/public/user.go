package public_dto

import (
	"strings"
	"time"

	"flower-backend/models"
	"flower-backend/utils"
)

type PublicUserDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type AuthOwnerUserDTO struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	Provider  string    `json:"provider"`
}

func ToPublicUser(user *models.User) PublicUserDTO {
	if user == nil {
		return PublicUserDTO{}
	}

	return PublicUserDTO{
		ID:       user.ID,
		Username: utils.SanitizeString(user.Username),
		Avatar:   utils.SanitizeURL(user.Avatar),
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

func ToAuthOwnerUser(user *models.User) AuthOwnerUserDTO {
	if user == nil {
		return AuthOwnerUserDTO{}
	}

	return AuthOwnerUserDTO{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
		Provider:  user.Provider,
	}
}
