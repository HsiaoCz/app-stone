package types

import "gorm.io/gorm"

type Addresses struct {
	gorm.Model
	AddressID     string `gorm:"column:address_id;" json:"address_id"`
	UserID        string `gorm:"column:user_id;" json:"user_id"`
	RecipientName string `gorm:"column:recipient_name;" json:"recipient_name"`
	AddressLine   string `gorm:"column:addrss_line;" json:"column:address_line"`
	City          string `gorm:"column:city;" json:"city"`
	PotalCode     int    `gorm:"column:potal_code;" json:"potal_code"`
	PhoneNumber   string `gorm:"column:phone_number;" json:"phone_number"`
}
