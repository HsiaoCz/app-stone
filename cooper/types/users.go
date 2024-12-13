package types

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	UserID string `gorm:"column:user_id" json:"user_id"`
	Email  string `gorm:"column:email" json:"email"`
}

type UserInfo struct {
	UserID string
	Email  string
	Role   bool
}
