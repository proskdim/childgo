package handler

import (
	model "childgo/app/models"
	"childgo/app/models/repo"
	"childgo/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

	u := &model.User{}

	if err := repo.FindUser(u, "email", email).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	return ctx.JSON(ProfileResponse{
		Email: u.Email,
	})
}
