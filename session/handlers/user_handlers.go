package handlers

import (
	"net/http"
	"time"

	"github.com/HsiaoCz/app-stone/session/storage"
	"github.com/HsiaoCz/app-stone/session/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandlers struct {
	user    storage.UserStoreInter
	session storage.SessionStoreInter
}

func UserHandlersInit(user storage.UserStoreInter, session storage.SessionStoreInter) *UserHandlers {
	return &UserHandlers{
		user:    user,
		session: session,
	}
}

func (u *UserHandlers) HandleCreateUser(c *fiber.Ctx) error {
	var user_create_params types.CreateUserParams
	if err := c.BodyParser(&user_create_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	msg := user_create_params.Validate()
	if len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}
	user := types.NewUserFromParams(user_create_params)
	userr, err := u.user.CreateUser(c.Context(), user)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	sessionu := &types.Sessions{
		Token:     uuid.New().String(),
		UserID:    userr.ID.String(),
		IpAddress: c.IP(),
		UserAgent: string(c.Context().UserAgent()),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	}
	sessionr, err := u.session.CreateSession(c.Context(), sessionu)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    sessionr.Token,
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
	})
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "create user success",
		"user":    userr,
		"session": sessionr,
	})
}

func (u *UserHandlers) HandleGetUserByID(c *fiber.Ctx) error {
	user_id, err := primitive.ObjectIDFromHex(c.Params("uid"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	user, err := u.user.GetUserByID(c.Context(), user_id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "get user successful",
		"user":    user,
	})
}

func (u *UserHandlers) HandleUserLoginOut(c *fiber.Ctx) error {
	return nil
}

func (u *UserHandlers) HandleUserLogin(c *fiber.Ctx) error {
	return nil
}
