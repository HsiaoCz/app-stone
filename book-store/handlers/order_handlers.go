package handlers

import "net/http"

type OrderHandlers struct{}

func (o *OrderHandlers) CreateOrder(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (o *OrderHandlers) HandleGetOrdersByUserID(w http.ResponseWriter, r *http.Request) error {
	return nil
}
