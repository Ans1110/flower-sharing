package services

import (
	"flower-backend/database"
	"flower-backend/models"
	"strconv"

	"gorm.io/gorm"
)

type PostWithAuthor struct {
	models.Post
	AuthorUsername string `json:"author_username"`
}

type PostService struct{}

func NewPostService() *PostService {
	return &PostService{}
}

// GetAllPostsWithAuthor retrieves all posts with author usernames
func (s *PostService) GetAllPostsWithAuthor() ([]PostWithAuthor, error) {
	var posts []PostWithAuthor

	err := database.DB.Table("posts").
		Select("posts.*, COALESCE(users.username, 'Unknown User') as author_username").
		Joins("LEFT JOIN users ON posts.author_id = users.id").
		Order("posts.created_at DESC").
		Scan(&posts).Error

	return posts, err
}

// GetPostWithAuthorByID retrieves a single post with author username by ID
func (s *PostService) GetPostWithAuthorByID(id string) (PostWithAuthor, error) {
	var post PostWithAuthor

	err := database.DB.Table("posts").
		Select("posts.*, COALESCE(users.username, 'Unknown User') as author_username").
		Joins("LEFT JOIN users ON posts.author_id = users.id").
		Where("posts.id = ?", id).
		First(&post).Error

	return post, err
}

// CreatePost creates a new post
func (s *PostService) CreatePost(title, content, imageURL string, authorID uint) (*models.Post, error) {
	// Get author username
	var author models.User
	if err := database.DB.Select("username").First(&author, authorID).Error; err != nil {
		return nil, err
	}

	post := &models.Post{
		Title:      title,
		Content:    content,
		ImageURL:   imageURL,
		AuthorID:   authorID,
		AuthorName: author.Username,
	}

	if err := database.DB.Create(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

// GetPostByID retrieves a post by ID
func (s *PostService) GetPostByID(id string) (*models.Post, error) {
	var post models.Post
	err := database.DB.First(&post, id).Error
	return &post, err
}

// UpdatePost updates an existing post
func (s *PostService) UpdatePost(id string, title, content, imageURL string) (*models.Post, error) {
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		return nil, err
	}

	post.Title = title
	post.Content = content
	post.ImageURL = imageURL

	if err := database.DB.Save(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

// DeletePost deletes a post by ID
func (s *PostService) DeletePost(id string) error {
	return database.DB.Delete(&models.Post{}, id).Error
}

// LikePost increments the like count for a post
func (s *PostService) LikePost(id string) (*models.Post, error) {
	// First check if post exists
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		return nil, err
	}

	// Increment likes
	if err := database.DB.Model(&post).Update("likes", gorm.Expr("likes + 1")).Error; err != nil {
		return nil, err
	}

	// Return updated post
	if err := database.DB.First(&post, id).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

// UnlikePost decrements the like count for a post
func (s *PostService) UnlikePost(id string) (*models.Post, error) {
	// First check if post exists
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		return nil, err
	}

	// Decrement likes
	if err := database.DB.Model(&post).Update("likes", gorm.Expr("likes - 1")).Error; err != nil {
		return nil, err
	}

	// Return updated post
	if err := database.DB.First(&post, id).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

// CheckPostOwnership checks if a user owns a post
func (s *PostService) CheckPostOwnership(postID string, userID uint) (bool, error) {
	post, err := s.GetPostByID(postID)
	if err != nil {
		return false, err
	}
	return post.AuthorID == userID, nil
}

// ValidatePostID validates if the post ID is valid
func (s *PostService) ValidatePostID(id string) error {
	_, err := strconv.ParseUint(id, 10, 32)
	return err
}
