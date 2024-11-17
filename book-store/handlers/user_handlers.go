package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/app-stone/book-store/data"
	"github.com/HsiaoCz/app-stone/book-store/handlers/middlewares"
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

func (u *UserHandlers) HandleUserLogin(w http.ResponseWriter, r *http.Request) error {
	var user_login_params types.UserLoginParams
	if err := json.NewDecoder(r.Body).Decode(&user_login_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	user, err := u.user.GetUserByEmailAndPassword(r.Context(), &user_login_params)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, "there is no record,please signup")
	}
	token, err := middlewares.GenToken(user.UserID, user.Email, user.Role)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, H{
		"status":  http.StatusOK,
		"message": "login success",
		"token":   token,
		"data":    user,
	})
}

func (u *UserHandlers) HandleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	userInfo, ok := r.Context().Value(types.CtxUserInfoKey).(types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "please login")
	}
	user, err := u.user.GetUserByID(r.Context(), userInfo.UserID)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, H{
		"status":  http.StatusOK,
		"message": "get user success",
		"user":    user,
	})
}

func (u *UserHandlers) HandleDeleteUserByID(w http.ResponseWriter, r *http.Request) error {
	userInfo, ok := r.Context().Value(types.CtxUserInfoKey).(types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "please login")
	}
	if err := u.user.DeleteUserByID(r.Context(), userInfo.UserID); err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, H{
		"status":  http.StatusOK,
		"message": "delete user success",
	})
}
