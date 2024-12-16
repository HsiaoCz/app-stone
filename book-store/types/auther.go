package types

import "gorm.io/gorm"

type Authers struct {
	gorm.Model
	AutherID   string `gorm:"column:auther_id;" json:"auther"`
	AutherName string `gorm:"column:auther_name" json:"auther_name"`
	Bio        string `gorm:"column:bio;" json:"bio"`
	Picture    string `gorm:"column:picture" json:"picture"`
}
