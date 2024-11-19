package types

import "gorm.io/gorm"

type Authers struct {
	gorm.Model
	AutherID string `gorm:"column:auther_id;" json:"auther"`
	Image    string `gorm:"column:image;" json:"image"`
}
