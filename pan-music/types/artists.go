package types

import "gorm.io/gorm"


type Artists struct{
	gorm.Model
	ArtistsID string `gorm:""`
}