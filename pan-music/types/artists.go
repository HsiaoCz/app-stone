package types

import "gorm.io/gorm"

type Artists struct {
	gorm.Model
	ArtistsID string `gorm:"column:artist_id" json:"artist_id"`
	Name      string `gorm:"column:name" json:"name"`
	Bio       string `gorm:"column:bio" json:"bio"`
}
