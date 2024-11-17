package types

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	BookID       string  `gorm:"column:book_id;" json:"book_id"`
	Title        string  `gorm:"column:title;" json:"title"`
	Auther       string  `gorm:"column:auther;" json:"auther"`
	Price        float64 `gorm:"column:price;" json:"price"`
	CategoryID   string  `gorm:"column:category_id;" json:"category_id"`
	Descriptions string  `gorm:"column:descriptions;" json:"descriptions"`
	Stock        int     `gorm:"column:stock;" json:"stock"`
	CoverImage   string  `gorm:"column:cover_image;" json:"cover_image"`
}
