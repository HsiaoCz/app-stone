package types

import "gorm.io/gorm"

type Albums struct {
	gorm.Model
	AlbumID     string `gorm:"column:album_id" json:"album"`
	Title       string `gorm:"column:title" json:"title"`
	ArtistID    string `gorm:"column:artist_id" json:"artist_id"`
	ReleaseDate string `gorm:"column:release_date" json:"release_date"`
	CoverImage  string `gorm:"column:cover_image" json:"cover_image"`
}
