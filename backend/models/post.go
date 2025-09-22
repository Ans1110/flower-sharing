package models

import "time"

type Post struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Content   string
	ImageURL  string
	AuthorID  uint
	CreatedAt time.Time
}
