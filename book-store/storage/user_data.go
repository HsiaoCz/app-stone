package storage

import (
	"context"

	"github.com/HsiaoCz/app-stone/book-store/pkgs"
	"github.com/HsiaoCz/app-stone/book-store/types"
	"gorm.io/gorm"
)

type UserDataInter interface {
	CreateUser(context.Context, *types.Users) (*types.Users, error)
	GetUserByEmailAndPassword(context.Context, *types.UserLoginParams) (*types.Users, error)
	GetUserByID(context.Context, string) (*types.Users, error)
	DeleteUserByID(context.Context, string) error
	GetUsersByUsername(context.Context, string) ([]*types.Users, error)
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

func (u *UserData) GetUserByID(ctx context.Context, user_id string) (*types.Users, error) {
	var user types.Users
	tx := u.db.Debug().WithContext(ctx).Model(&types.Users{}).Where("user_id = ?", user_id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (u *UserData) DeleteUserByID(ctx context.Context, user_id string) error {
	tx := u.db.Debug().WithContext(ctx).Model(&types.Users{}).Delete("user_id = ?", user_id)
	return tx.Error
}

func (u *UserData) GetUsersByUsername(ctx context.Context, username string) ([]*types.Users, error) {
	var users []*types.Users
	tx := u.db.Debug().WithContext(ctx).Model(&types.Users{}).Where("username = ?", username).Find(&users)
	return users, tx.Error
}
