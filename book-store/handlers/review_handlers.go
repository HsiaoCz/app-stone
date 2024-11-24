package handlers

import "net/http"

type ReviewHandlers struct{}

func (v *ReviewHandlers) HandleCreateReview(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (v *ReviewHandlers) HandleGetReviewByBookID(w http.ResponseWriter, r *http.Request) error {
	return nil
}
