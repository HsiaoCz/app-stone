package data

import (
	"context"

	"github.com/HsiaoCz/app-stone/book-store/pkgs"
	"github.com/HsiaoCz/app-stone/book-store/types"
	"gorm.io/gorm"
)

type UserDataInter interface {
	CreateUser(context.Context, *types.Users) (*types.Users, error)
	GetUserByEmailAndPassword(context.Context, *types.UserLoginParams) (*types.Users, error)
}

type UserData struct {
	db *gorm.DB
}

func UserDataInit(db *gorm.DB) *UserData {
	return &UserData{
		db: db,
	}
}

func (u *UserData) CreateUser(ctx context.Context, user *types.Users) (*types.Users, error) {
	tx := u.db.Debug().WithContext(ctx).Model(&types.Users{}).Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *UserData) GetUserByEmailAndPassword(ctx context.Context, params *types.UserLoginParams) (*types.Users, error) {
	var user types.Users
	tx := u.db.Debug().WithContext(ctx).Model(&types.Users{}).Where("email = ? AND hash_password = ?", params.Email, pkgs.EncryPassword(params.Password)).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
