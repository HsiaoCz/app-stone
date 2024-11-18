package handlers

import "net/http"

type AddressHandlers struct{}

func (a *AddressHandlers) HandleCreateAddress(w http.ResponseWriter, r *http.Request) error {
	return nil
}
