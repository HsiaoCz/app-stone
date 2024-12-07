package storage

import (
	"context"

	"github.com/HsiaoCz/app-stone/ebuy/types"
)


type UserStorer interface{
	CreateUser(context.Context,*types.Users)(*types.Users,error)
}