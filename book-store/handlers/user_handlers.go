package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/app-stone/book-store/data"
	"github.com/HsiaoCz/app-stone/book-store/types"
)

type UserHandlers struct {
	user data.UserDataInter
}

func UserHandlersInit(user data.UserDataInter) *UserHandlers {
	return &UserHandlers{
		user: user,
	}
}

func (u *UserHandlers) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	var create_user_params types.CreateUserParams
	if err := json.NewDecoder(r.Body).Decode(&create_user_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	user := types.CreateUserFromParams(create_user_params)
	userReturn, err := u.user.CreateUser(r.Context(), user)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, H{
		"status":  http.StatusOK,
		"message": "create user success",
		"data":    userReturn,
	})
}
