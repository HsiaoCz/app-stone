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
	return nil
}
