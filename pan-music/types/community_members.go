package types

import "gorm.io/gorm"

type Community_Members struct {
	gorm.Model
	MemberID string `gorm:"column:member_id"`
}
