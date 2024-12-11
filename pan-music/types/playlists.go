package types

import "gorm.io/gorm"

type Playlists struct {
	gorm.Model
	PlayListID   string `gorm:"column:playlist_id" json:"playlist_id"`
	UserID       string `gorm:"column:user_id" json:"user_id"`
	Title        string `gorm:"column:title" json:"title"`
	Descriptions string `gorm:"column:descriptions" json:"descriptions"`
}
