package admin_dto

import (
	"flower-backend/models"
	"strings"
	"time"
)

type UserAdminDTO struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	Role      string    `json:"role"`
	Posts     int       `json:"posts"`
	Likes     int       `json:"likes"`
	Followers int       `json:"followers"`
	Following int       `json:"following"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToUserAdminDTO(user *models.User) UserAdminDTO {
	if user == nil {
		return UserAdminDTO{}
	}

	return UserAdminDTO{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Role:      user.Role,
		Posts:     len(user.Posts),
		Likes:     len(user.Likes),
		Followers: len(user.Followers),
		Following: len(user.Following),
	}
}

func ToUserAdminDTOs(users []models.User) []UserAdminDTO {
	result := make([]UserAdminDTO, 0, len(users))
	for i := range users {
		result = append(result, ToUserAdminDTO(&users[i]))
	}
	return result
}

func EnsureUserAdminSelectFields(fields []string) []string {
	required := map[string]struct{}{
		"id":        {},
		"username":  {},
		"email":     {},
		"avatar":    {},
		"role":      {},
		"posts":     {},
		"likes":     {},
		"followers": {},
		"following": {},
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
