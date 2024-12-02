package handlers

import (
	"github.com/HsiaoCz/app-stone/session/storage"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	user    storage.UserStoreInter
	session storage.SessionStoreInter
}

func UserHandlersInit(user storage.UserStoreInter,session storage.SessionStoreInter) *UserHandlers {
	return &UserHandlers{
		user: user,
		session: session,
	}
}

func (u *UserHandlers) HandleCreateUser(c *fiber.Ctx) error {
	return nil
}
