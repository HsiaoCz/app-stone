package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/app-stone/book-store/storage"
	"github.com/HsiaoCz/app-stone/book-store/types"
)

type ReviewHandlers struct {
	dr storage.ReviewDataInter
}

func ReviewHandlersInit(dr storage.ReviewDataInter) *ReviewHandlers {
	return &ReviewHandlers{
		dr: dr,
	}
}

func (v *ReviewHandlers) HandleCreateReview(w http.ResponseWriter, r *http.Request) error {
	userInfo, ok := r.Context().Value(types.CtxUserInfoKey).(*types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "please login to do this shit......")
	}
	var params types.CreateReviewParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	review := types.CreateReviewFromParams(params)
	review.UserID = userInfo.UserID
	reviewR, err := v.dr.CreateReview(r.Context(), review)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, reviewR)
}

func (v *ReviewHandlers) HandleGetReviewByBookID(w http.ResponseWriter, r *http.Request) error {
	book_id := r.URL.Query().Get("bid")
	reviews, err := v.dr.GetReviewByBookID(r.Context(), book_id)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	return WriteJson(w, http.StatusOK, reviews)
}

func (v *ReviewHandlers) HandleGetReviewByID(w http.ResponseWriter, r *http.Request) error {
	review_id := r.URL.Query().Get("rid")
	review, err := v.dr.GetReviewByID(r.Context(), review_id)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	return WriteJson(w, http.StatusOK, review)
}

func (v *ReviewHandlers) HandleDeleteReviewByID(w http.ResponseWriter, r *http.Request) error {
	review_id := r.PathValue("rid")
	if err := v.dr.DeleteReview(r.Context(), review_id); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	return WriteJson(w, http.StatusOK, H{
		"status":  http.StatusOK,
		"message": "delete review success",
	})
}
