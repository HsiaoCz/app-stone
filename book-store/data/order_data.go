package data

import (
	"context"

	"github.com/HsiaoCz/app-stone/book-store/types"
)

type OrderDataInter interface {
	CreateOrder(context.Context, *types.Orders) (*types.Orders, error)
}


