package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/app-stone/book-store/data"
	"github.com/HsiaoCz/app-stone/book-store/types"
)

type BookHandlers struct {
	book data.BookDataInter
}

func BookHandlersInit(book data.BookDataInter) *BookHandlers {
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

	return nil
}
