package data

import (
	"context"

	"github.com/HsiaoCz/app-stone/book-store/types"
)

type ReviewDataInter interface {
	CreateReview(context.Context, *types.Reviews) (*types.Reviews, error)
}
