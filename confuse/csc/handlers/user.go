package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HsiaoCz/app-stone/confuse/csc/storage"
	"github.com/HsiaoCz/app-stone/confuse/csc/types"
	"github.com/google/uuid"
)

type UserHandlers struct {
	usr storage.UserStorer
	ser storage.SessionStorer
}

func UserHandlersInit(usr storage.UserStorer, ser storage.SessionStorer) *UserHandlers {
	return &UserHandlers{
		usr: usr,
		ser: ser,
	}
}

func (u *UserHandlers) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	var create_user_params types.CreateUserParams
	if err := json.NewDecoder(r.Body).Decode(&create_user_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	user, err := u.usr.CreateUser(r.Context(), types.CreateUserFromParams(create_user_params))
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
	var user_login_params types.UserLoginParams
	if err := json.NewDecoder(r.Body).Decode(&user_login_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	user, err := u.usr.GetUserByEmailAndPassword(r.Context(), user_login_params)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	session := &types.Sessions{
		Token:     uuid.New().String(),
		UserID:    user.UserID,
		UserAgent: r.UserAgent(),
		IpAddress: r.RemoteAddr,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	}
	if err = u.ser.CreateSession(r.Context(), session); err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    session.Token,
		Expires:  session.ExpiresAt,
		HttpOnly: true,
	})
	return WriteJson(w, http.StatusOK, H{
		"status":  http.StatusOK,
		"message": "user login success",
		"user":    user,
	})
}

func (u *UserHandlers) HandleUserLogout(w http.ResponseWriter, r *http.Request) error {
	session, ok := r.Context().Value(types.UserSessionKey).(*types.Sessions)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "please login")
	}
	if err := u.ser.DeleteSessionByToken(r.Context(), session.Token); err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, H{
		"status":  http.StatusOK,
		"message": "delete user success",
		"user_id": session.UserID,
	})
}
