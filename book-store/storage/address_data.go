package storage

import (
	"context"

	"github.com/HsiaoCz/app-stone/book-store/types"
	"gorm.io/gorm"
)

type AddressDataInter interface {
	CreateAddress(context.Context, *types.Addresses) (*types.Addresses, error)
}

type AddressData struct {
	db *gorm.DB
}

func AddressDataInit(db *gorm.DB) *AddressData {
	return &AddressData{
		db: db,
	}
}

func (a *AddressData) CreateAddress(ctx context.Context, address *types.Addresses) (*types.Addresses, error) {
	tx := a.db.Debug().WithContext(ctx).Model(&types.Addresses{}).Create(address)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return address, nil
}
