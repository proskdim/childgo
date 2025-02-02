package handler

import (
	"childgo/config"
	"childgo/database"
	"childgo/model/user"
	jwtUtil "childgo/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
)

type ProfileResponse struct {
	Email string `json:"email"`
}

var (
	ErrUserNotFound = errors.New("user not found")
	ErrJwtContext   = errors.New("wrong type of JWT token in context")
)

func Profile(ctx *fiber.Ctx) error {
	token, ok := ctx.Locals(config.ContextKeyUser).(*jwt.Token)

	if !ok {
		return ctx.Status(fiber.StatusBadRequest).SendString(ErrJwtContext.Error())
	}

	email, err := jwtUtil.FindEmailFromToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	dbUser, err := user.FindByEmail(database.DBConn, email)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return ctx.JSON(ProfileResponse{
		Email: dbUser.Email,
	})
}
