package handlers

import "net/http"

type OrderHandlers struct{}

func (o *OrderHandlers) CreateOrder(w http.ResponseWriter, r *http.Request) error {
	return nil
}
