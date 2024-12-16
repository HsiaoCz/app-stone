package storage

import (
	"context"

	"github.com/HsiaoCz/app-stone/cooper/types"
)

type SessionStoreInter interface {
	CreateSession(context.Context, *types.Sessions) error
}
