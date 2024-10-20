package types

import "gorm.io/gorm"

type Users struct {
	gorm.Model
}

type UserInfo struct {
	UserID string
	Email  string
	Role   bool
}
