package types

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	OrderID     string  `gorm:"column:order_id;" json:"order_id"`
	UserID      string  `gorm:"column:user_id;" json:"user_id"`
	OrderStatus string  `gorm:"column:order_status;" json:"order_status"`
	TotalAmount float64 `gorm:"column:total_amount;" json:"total_amount"`
}
