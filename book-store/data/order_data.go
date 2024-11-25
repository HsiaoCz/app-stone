package data

import (
	"context"

	"github.com/HsiaoCz/app-stone/book-store/types"
	"gorm.io/gorm"
)

type OrderDataInter interface {
	CreateOrder(context.Context, *types.Orders) (*types.Orders, error)
}

type OrderData struct {
	db *gorm.DB
}

func OrderDataInit(db *gorm.DB) *OrderData {
	return &OrderData{
		db: db,
	}
}

func (o *OrderData) CreateOrder(ctx context.Context, order *types.Orders) (*types.Orders, error) {
	tx := o.db.Debug().WithContext(ctx).Model(&types.Orders{}).Create(order)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return order, nil
}
