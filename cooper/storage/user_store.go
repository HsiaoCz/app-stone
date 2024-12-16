package storage

import (
	"context"

	"github.com/HsiaoCz/app-stone/cooper/types"
)

type UserStoreInter interface {
	CreateUser(context.Context, *types.Users) (*types.Users, error)
}
