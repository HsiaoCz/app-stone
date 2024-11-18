package types

import "gorm.io/gorm"

type Records struct {
	gorm.Model
	RecordID string `gorm:"column:record_id;" json:"record_id"`
	UserID   string `gorm:"column:user_id;" json:"user_id"`
	BookID   string `gorm:"column:book_id;" json:"book_id"`
	TypeName string `gorm:"column:type_name;" json:"type_name"`
	Device   string `gorm:"column:device;" json:"device"`
}
