package middlewares

import (
	"net/http"

	"github.com/HsiaoCz/app-stone/session/handlers"
	"github.com/HsiaoCz/app-stone/session/storage"
	"github.com/HsiaoCz/app-stone/session/types"
	"github.com/gofiber/fiber/v2"
)

type AuthMdw struct {
	session storage.SessionStoreInter
}

func AuthMdwInit(session storage.SessionStoreInter) *AuthMdw {
	return &AuthMdw{
		session: session,
	}
}

func (a *AuthMdw) AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return handlers.ErrorMessage(http.StatusNonAuthoritativeInfo, "please login")
	}
	session, err := a.session.GetSessionByToken(c.Context(), token)
	if err != nil {
		return handlers.ErrorMessage(http.StatusNonAuthoritativeInfo, "please login")
	}
	c.Locals(types.CtxUserInfoKey, session)
	return c.Next()
}
