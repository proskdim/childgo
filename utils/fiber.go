package utils

import (
	model "childgo/app/models"
	"childgo/config"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// ParseBody is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBody(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

// ParseID is helper function for parsing query param id
func ParseID(ctx *fiber.Ctx) (*uuid.UUID, *fiber.Error) {
	uuid, err := uuid.Parse(ctx.Params("id", ""))

	if err != nil {
		return nil, fiber.ErrBadRequest
	}

	return &uuid, nil
}

// GetUser is helper function for getting authenticated user's id
func GetUser(ctx *fiber.Ctx) *model.User {
	id, _ := ctx.Locals(config.ContextJwtUser).(*model.User)
	return id
}
