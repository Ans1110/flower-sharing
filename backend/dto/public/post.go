package public_dto

import (
	"flower-backend/models"
	"flower-backend/utils"
	"time"
)

type PublicPostDTO struct {
	ID        uint          `json:"id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	ImageURL  string        `json:"image_url"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Author    PublicUserDTO `json:"author"`
	Likes     int           `json:"likes_count"`
}

func ToPublicPost(post *models.Post) PublicPostDTO {
	if post == nil {
		return PublicPostDTO{}
	}

	return PublicPostDTO{
		ID:        post.ID,
		Title:     utils.SanitizeString(post.Title),
		Content:   utils.SanitizeHTML(post.Content),
		ImageURL:  utils.SanitizeURL(post.ImageURL),
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Author:    ToPublicUser(&post.User),
		Likes:     len(post.Likes),
	}
}

func ToPublicPosts(posts []models.Post) []PublicPostDTO {
	result := make([]PublicPostDTO, 0, len(posts))
	for i := range posts {
		result = append(result, ToPublicPost(&posts[i]))
	}
	return result
}
