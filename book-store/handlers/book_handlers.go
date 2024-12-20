package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/app-stone/book-store/storage"
	"github.com/HsiaoCz/app-stone/book-store/types"
)

type BookHandlers struct {
	book storage.BookDataInter
}

func BookHandlersInit(book storage.BookDataInter) *BookHandlers {
	return &BookHandlers{
		book: book,
	}
}

func (b *BookHandlers) HandleCreateBook(w http.ResponseWriter, r *http.Request) error {
	userInfo, ok := r.Context().Value(types.CtxUserInfoKey).(types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "please login")
	}
	if !userInfo.Role {
		return ErrorMessage(http.StatusNotAcceptable, "your are not admin so you have no rights to do this")
	}
	var book types.Books
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	bookReturn, err := b.book.CreateBook(r.Context(), &book)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, H{
		"status":  http.StatusOK,
		"message": "create book success",
		"book":    bookReturn,
	})
}

func (b *BookHandlers) HandleGetBookByID(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (b *BookHandlers) HandleDeleteBookByID(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (b *BookHandlers) HandleUpdateBook(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (b *BookHandlers) HandleGetBookByAuther(w http.ResponseWriter, r *http.Request) error {
	return nil
}
