package handlers

import "net/http"

type BookHandlers struct{}

func (b *BookHandlers) HandleCreateBook(w http.ResponseWriter, r *http.Request) error {
	return nil
}
