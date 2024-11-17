package handlers

import (
	"net/http"

	"github.com/HsiaoCz/app-stone/pan-music/data"
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
	return nil
}
