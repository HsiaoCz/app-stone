package handlers

import "net/http"

type HomeHandlers struct{}

func HomeHandlersInit() *HomeHandlers {
	return &HomeHandlers{}
}

func (h *HomeHandlers) HandleHomePage(w http.ResponseWriter, r *http.Request) error {
	return nil
}
