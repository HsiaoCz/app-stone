package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reviews struct {
	gorm.Model
	ReviewID     string  `gorm:"column:review_id;" json:"review_id"`
	BaseReviewID string  `gorm:"column:base_review_id;" json:"base_review_id"`
	UserID       string  `gorm:"column:user_id;" json:"user_id"`
	BookID       string  `gorm:"column:book_id;" json:"book_id"`
	Rating       float64 `gorm:"column:rating;" json:"rating"`
	Comment      string  `gorm:"column:comment;" json:"comment"`
}

type CreateReviewParams struct {
	BaseReviewID string  `json:"base_review_id"`
	BookID       string  `json:"book_id"`
	Rating       float64 `json:"rating"`
	Comment      string  `json:"comment"`
}

func CreateReviewFromParams(params CreateReviewParams) *Reviews {
	return &Reviews{
		ReviewID:     uuid.New().String(),
		BaseReviewID: params.BaseReviewID,
		BookID:       params.BookID,
		Rating:       params.Rating,
	}
}
