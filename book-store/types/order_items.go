package types

import "gorm.io/gorm"

type Order_Items struct {
	gorm.Model
	OrderItemID string  `gorm:"column:order_item_id;" json:"order_item_id"`
	OrderID     string  `gorm:"column:order_id;" json:"order_id"`
	BookID      string  `gorm:"column:book_id;" json:"book_id"`
	Quantity    int     `gorm:"column:quantity;" json:"quantity"`
	Price       float64 `gorm:"column:price;" json:"price"`
}
