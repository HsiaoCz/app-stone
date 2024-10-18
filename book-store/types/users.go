package types

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	UserID       string `gorm:"column:user_id;" json:"user_id"`
	Username     string `gorm:"column:username;" json:"username"`
	Email        string `gorm:"column:email;" json:"email"`
	PasswordHash string `gorm:"column:password_hash;" json:"-"`
	Role         bool   `gorm:"column:role;" json:"role"`
}
