package types

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	UserID       string `gorm:"column:user_id" json:"user_id"`
	Email        string `gorm:"column:email" json:"email"`
	Username     string `gorm:"column:username" json:"username"`
	PasswordHash string `gorm:"column:password_hash" json:"-"`
	Role         bool   `gorm:"column:role" json:"role"`
}

type UserInfo struct {
	UserID string
	Email  string
	Role   bool
}
