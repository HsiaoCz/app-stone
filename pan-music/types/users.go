package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UserID       string `gorm:"column:user_id;" json:"user_id"`
	Username     string `gorm:"column:username;" json:"username"`
	Email        string `gorm:"column:email;" json:"email"`
	PasswordHash string `gorm:"column:password_hash;" json:"-"`
	Role         bool   `gorm:"column:role;" json:"role"`
}

type UserInfo struct {
	UserID string
	Email  string
	Role   bool
}

type CreateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     bool   `json:"role"`
}

func CreateUserFromParams(params CreateUserParams) *Users {
	return &Users{
		UserID:       uuid.New().String(),
		Username:     params.Username,
		Email:        params.Email,
		PasswordHash: params.Password,
		Role:         params.Role,
	}
}
