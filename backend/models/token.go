package models

import "time"

type Token struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Token     string    `gorm:"not null;unique" json:"token"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	ExpiresAt time.Time `gorm:"index" json:"expires_at"`
}
