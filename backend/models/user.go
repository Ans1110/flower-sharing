package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"not null;unique" json:"username"`
	Email        string    `gorm:"not null;unique" json:"email"`
	Password     string    `json:"-"`
	Avatar       string    `json:"avatar"`
	CreatedAt    time.Time `json:"created_at"`
	Role         string    `gorm:"default:user" json:"role"`
	Provider     string    `gorm:"default:local" json:"provider"`  // local, google, github
	ProviderID   string    `json:"provider_id"`                    // OAuth provider user ID
	ProviderData string    `gorm:"type:text" json:"provider_data"` // JSON data from OAuth provider
	Posts        []Post    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID" json:"posts"`
	Tokens       []Token   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID" json:"-"`
	Likes        []Post    `gorm:"many2many:post_likes" json:"likes"`
	Followers    []User    `gorm:"many2many:user_follows;joinForeignKey:following_id;joinReferences:follower_id" json:"followers"`
	Following    []User    `gorm:"many2many:user_follows;joinForeignKey:follower_id;joinReferences:following_id" json:"following"`
}
