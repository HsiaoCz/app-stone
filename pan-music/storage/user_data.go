package storage

import (
	"context"

	pkg "github.com/HsiaoCz/app-stone/pan-music/pkgs"
	"github.com/HsiaoCz/app-stone/pan-music/types"
	"gorm.io/gorm"
)

type UserDataInter interface {
	CreateUser(context.Context, *types.Users) (*types.Users, error)
	GetUserByEmailAndPassword(context.Context, *types.UserLogin) (*types.Users, error)
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

func (u *UserData) GetUserByEmailAndPassword(ctx context.Context, params *types.UserLogin) (*types.Users, error) {
	var user types.Users
	tx := u.db.Debug().WithContext(ctx).Model(&types.Users{}).Where("email = ? AND password = ?", params.Email, pkg.EncryPassword(params.Password)).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
