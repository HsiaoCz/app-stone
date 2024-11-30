package types

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	CommentID string `gorm:"column:comment_id;" json:"comment_id"`
}
