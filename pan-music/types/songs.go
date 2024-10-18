package types

import "gorm.io/gorm"

type Songs struct {
	gorm.Model
	SongID      string `gorm:"column:song_id;" json:"song_id"`
	Title       string `gorm:"column:title;" json:"title"`
	ArtistID    string `gorm:"column:artist_id;" json:"artist_id"`
	AlbumID     string `gorm:"column:album_id;" json:"album_id"`
	Genre       string `gorm:"column:genre;" json:"genre"`
	Durations   int    `gorm:"column:durations" json:"durations"`
	ReleaseDate string `gorm:"column:release_date;" json:"release_date"`
	Lyrics      string `gorm:"column:lyrics;" json:"lyrics"`
}
