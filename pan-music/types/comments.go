package types

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	CommentID string `gorm:"column:comment_id;" json:"comment_id"`
	UserID    string `gorm:"column:user_id" json:"user_id"`
	SongID    string `gorm:"column:song_id" json:"song_id"`
	Content   string `gorm:"column:content" json:"content"`
}
