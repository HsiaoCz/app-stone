package types

import (
	"time"

	"gorm.io/gorm"
)

type Sessions struct {
	gorm.Model
	Token     string    `gorm:"column:token;" json:"token"`
	UserID    string    `gorm:"column:user_id;" json:"user_id"`
	IpAddress string    `gorm:"column:ip_address;" json:"ip_address"`
	UserAgent string    `gorm:"column:user_agent;" json:"user_agent"`
	ExpiresAt time.Time `gorm:"column:expires_at;" json:"expires_at"`
}
