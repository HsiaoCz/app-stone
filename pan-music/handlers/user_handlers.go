package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/app-stone/pan-music/data"
	"github.com/HsiaoCz/app-stone/pan-music/handlers/middlewares"
	"github.com/HsiaoCz/app-stone/pan-music/types"
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
	user, err := u.user.CreateUser(r.Context(), types.CreateUserFromParams(create_user_params))
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, H{
		"status":  http.StatusOK,
		"message": "create user success",
		"user":    user,
	})
}

func (u *UserHandlers) HandleUserLogin(w http.ResponseWriter, r *http.Request) error {
	var user_login_params types.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&user_login_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	user, err := u.user.GetUserByEmailAndPassword(r.Context(), &user_login_params)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, "用户名或密码错误")
	}
	token, err := middlewares.GenToken(user.UserID, user.Email, user.Role)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, "有点儿坏了")
	}
	return WriteJson(w, http.StatusOK, H{
		"status":  http.StatusOK,
		"message": "user login success",
		"user":    user,
		"token":   token,
	})
}
