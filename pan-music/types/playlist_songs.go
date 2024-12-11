package types

import (
	"time"

	"gorm.io/gorm"
)

type Playlist_Songs struct {
	gorm.Model
	PlayListID string    `gorm:"column:playlist_id" json:"playlist_id"`
	SongID     string    `gorm:"column:song_id" json:"song_id"`
	AddedAt    time.Time `gorm:"column:added_at" json:"added_at"`
}
