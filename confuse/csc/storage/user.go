package storage

import (
	"context"

	"github.com/HsiaoCz/app-stone/confuse/csc/pkgs"
	"github.com/HsiaoCz/app-stone/confuse/csc/types"
	"gorm.io/gorm"
)

type UserStorer interface {
	CreateUser(context.Context, *types.Users) (*types.Users, error)
	GetUserByEmailAndPassword(context.Context, types.UserLoginParams) (*types.Users, error)
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

func (u *UserStore) GetUserByEmailAndPassword(ctx context.Context, params types.UserLoginParams) (*types.Users, error) {
	var user types.Users
	tx := u.db.Debug().WithContext(ctx).Model(&types.Users{}).Where("email = ? AND password_hash = ?", params.Email, pkgs.EncryPassword(params.Password)).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
