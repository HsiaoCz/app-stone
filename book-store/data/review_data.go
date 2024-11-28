package data

import (
	"context"

	"github.com/HsiaoCz/app-stone/book-store/types"
	"gorm.io/gorm"
)

type ReviewDataInter interface {
	CreateReview(context.Context, *types.Reviews) (*types.Reviews, error)
	GetReviewByBookID(context.Context, string) ([]*types.Reviews, error)
	DeleteReview(context.Context, string) error
	GetReviewByID(context.Context, string) (*types.Reviews, error)
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

func (r *ReviewData) GetReviewByBookID(ctx context.Context, book_id string) ([]*types.Reviews, error) {
	var reviews []*types.Reviews
	tx := r.db.Debug().WithContext(ctx).Model(&types.Reviews{}).Where("book_id = ?", book_id).Find(&reviews)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return reviews, nil
}

func (r *ReviewData) GetReviewByID(ctx context.Context, review_id string) (*types.Reviews, error) {
	var review types.Reviews
	tx := r.db.Debug().WithContext(ctx).Model(&types.Reviews{}).Where("review_id = ?", review_id).First(&review)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &review, nil
}

func (r *RecordData)DeleteReview(ctx context.Context,review_id string)error{
	var review types.Reviews
	tx:=r.db.Debug().WithContext(ctx).Model(&types.Reviews{}).Where("review_id = ?",review_id).Delete(review)
	return tx.Error
}