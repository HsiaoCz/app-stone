package data

import (
	"context"

	"github.com/HsiaoCz/app-stone/book-store/types"
	"gorm.io/gorm"
)

type RecordDataInter interface {
	CreateRecord(context.Context, *types.Records) (*types.Records, error)
}

type RecordData struct {
	db *gorm.DB
}

func RecordDataInit(db *gorm.DB) *RecordData {
	return &RecordData{
		db: db,
	}
}

func (r *RecordData) CreateRecord(ctx context.Context, record *types.Records) (*types.Records, error) {
	tx := r.db.Debug().WithContext(ctx).Model(&types.Records{}).Create(record)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return record, nil
}
