package handlers

import (
	"net/http"

	"github.com/HsiaoCz/app-stone/confuse/csc/storage"
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
	return nil
}

func (u *UserHandlers) HandleUserLogin(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserHandlers) HandleUserLogout(w http.ResponseWriter, r *http.Request) error {
	return nil
}
