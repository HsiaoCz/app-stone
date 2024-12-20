package storage

import (
	"context"

	"github.com/HsiaoCz/app-stone/book-store/types"
	"gorm.io/gorm"
)

type OrderDataInter interface {
	CreateOrder(context.Context, *types.Orders) (*types.Orders, error)
	DeleteOrder(context.Context, string) error
	UpdateOrder(context.Context, string, string) (*types.Orders, error)
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

func (o *OrderData) DeleteOrder(ctx context.Context, order_id string) error {
	var order types.Orders
	tx := o.db.Debug().WithContext(ctx).Model(&types.Orders{}).Where("order_id = ?", order_id).Delete(&order)
	return tx.Error
}

func (o *OrderData) UpdateOrder(ctx context.Context, order_id string, order_status string) (*types.Orders, error) {
	var order types.Orders
	tx := o.db.Debug().WithContext(ctx).Model(&types.Orders{}).Where("order_id = ?", order_id).Update("order_status", order_status).First(&order)
	return &order, tx.Error
}
