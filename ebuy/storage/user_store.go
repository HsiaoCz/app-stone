package storage

import (
	"context"

	"github.com/HsiaoCz/app-stone/ebuy/types"
	"gorm.io/gorm"
)

type UserStorer interface {
	CreateUser(context.Context, *types.Users) (*types.Users, error)
}

type UserStore struct {
	db *gorm.DB
}

func UserStoreInit(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (u *UserStore) CreateUser(ctx context.Context, user *types.Users) (*types.Users, error) {
	tx := u.db.Debug().WithContext(ctx).Model(&types.Users{}).Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
