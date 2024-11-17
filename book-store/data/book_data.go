package data

import (
	"context"

	"github.com/HsiaoCz/app-stone/book-store/types"
	"gorm.io/gorm"
)

type BookDataInter interface {
	CreateBook(context.Context, *types.Books) (*types.Books, error)
}

type BookData struct {
	db *gorm.DB
}

func BookDataInit(db *gorm.DB) *BookData {
	return &BookData{
		db: db,
	}
}

func (b *BookData) CreateBook(ctx context.Context, book *types.Books) (*types.Books, error) {
	tx := b.db.Debug().WithContext(ctx).Model(&types.Books{}).Create(book)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return book, nil
}
