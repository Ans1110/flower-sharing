package models

type Token struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Token  string `gorm:"not null;unique;required" json:"token"`
	UserID uint   `gorm:"not null;required" json:"user_id"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID" json:"user"`
}
