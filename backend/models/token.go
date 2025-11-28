package models

type Token struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Token  string `gorm:"not null;unique" json:"token"`
	UserID uint   `gorm:"not null" json:"user_id"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
}
