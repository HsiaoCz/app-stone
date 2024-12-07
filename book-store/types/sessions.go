package types

import (
	"time"

	"gorm.io/gorm"
)

type Sessions struct {
	gorm.Model
	Token     string
	UserID    string
	UserAgent string
	IpAddress string
	ExpiresAt time.Time
}
