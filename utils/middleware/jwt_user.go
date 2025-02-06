package middleware

import (
	model "childgo/app/models"
	"childgo/app/models/repo"
	"childgo/config"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// fetch user before api handlers
func JwtUserMiddleware(ctx *fiber.Ctx) error {
	token, ok := ctx.Locals(config.ContextKeyUser).(*jwt.Token)

	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "wrong type of JWT token in context"})
	}

	user, err := fetchUser(token)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	ctx.Locals(config.ContextJwtUser, user)

	return ctx.Next()
}

// fetch user from db by jwt token
func fetchUser(token *jwt.Token) (*model.User, error) {
	payload, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("failed to extraction token")
	}

	email := payload["sub"].(string)

	u := &model.User{}

	err := repo.FindUser(u, "email", email).Error

	if err != nil {
		return nil, errors.New("user by jwt token not found")
	}

	return u, nil
}
