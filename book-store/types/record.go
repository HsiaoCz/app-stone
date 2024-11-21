package types

import "gorm.io/gorm"

type Records struct {
	gorm.Model
	RecordID   string `gorm:"column:record_id;" json:"record_id"`
	UserID     string `gorm:"column:user_id;" json:"user_id"`
	BookID     string `gorm:"column:book_id;" json:"book_id"`
	BookName   string `gorm:"column:book_name;" json:"book_name"`
	CoverImage string `gorm:"column:cover_image;" json:"cover_image"`
	Auther     string `gorm:"column:auther;" json:"auther"`
	TypeName   string `gorm:"column:type_name;" json:"type_name"`
	Device     string `gorm:"column:device;" json:"device"`
}
