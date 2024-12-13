package types

import (
	"time"

	"gorm.io/gorm"
)

type Sessions struct {
	gorm.Model
	Token     string    `gorm:"column:token"`
	UserID    string    `gorm:"column:user_id"`
	UserAgent string    `gorm:"column:user_agent"`
	IpAddress string    `gorm:"column:ip_address"`
	ExpiresAt time.Time `gorm:"column:expires_at"`
}
