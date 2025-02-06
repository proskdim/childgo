package handler

import (
	"childgo/config"
	"childgo/config/database"
	"childgo/app/models/user"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
)

type ProfileResponse struct {
	Email string `json:"email"`
}

func Profile(ctx *fiber.Ctx) error {
	token, ok := ctx.Locals(config.ContextKeyUser).(*jwt.Token)

	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "wrong type of JWT token in context"})
	}

	payload, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to extraction token"})
	}

	email := payload["sub"].(string)

	dbUser, err := user.FindByEmail(database.DBConn, email)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	return ctx.JSON(ProfileResponse{
		Email: dbUser.Email,
	})
}
