package data

import (
	"context"

	"github.com/HsiaoCz/app-stone/book-store/types"
	"gorm.io/gorm"
)

type BookDataInter interface {
	CreateBook(context.Context, *types.Books) (*types.Books, error)
	GetBookByID(context.Context, string) (*types.Books, error)
	GetBookByAuther(context.Context, string) ([]*types.Books, error)
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

func (b *BookData) GetBookByID(ctx context.Context, book_id string) (*types.Books, error) {
	var book types.Books
	tx := b.db.Debug().WithContext(ctx).Model(&types.Books{}).Where("book_id = ?", book_id).First(&book)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &book, nil
}

func (b *BookData) GetBookByAuther(ctx context.Context, auther string) ([]*types.Books, error) {
	var books []*types.Books
	tx := b.db.Debug().WithContext(ctx).Model(&types.Books{}).Where("auther = ?", auther).Find(&books)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return books, nil
}
