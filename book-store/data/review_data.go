package data

import (
	"context"

	"github.com/HsiaoCz/app-stone/book-store/types"
	"gorm.io/gorm"
)

type ReviewDataInter interface {
	CreateReview(context.Context, *types.Reviews) (*types.Reviews, error)
}
type ReviewData struct {
	db *gorm.DB
}

func ReviewDataInit(db *gorm.DB) *RecordData {
	return &RecordData{
		db: db,
	}
}

func (r *ReviewData) CreateReview(ctx context.Context, review *types.Reviews) (*types.Reviews, error) {
	tx := r.db.Debug().WithContext(ctx).Model(&types.Reviews{}).Create(review)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return review, nil
}
