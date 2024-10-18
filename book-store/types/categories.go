package types

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	CategoryID string `gorm:"column:category_id;" json:"category_id"`
	CategoryName string `gorm:"column:category_name;" json:"category_name"`
}
