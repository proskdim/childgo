package utils

import (
	model "childgo/app/models"
	"childgo/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// ParseBody is helper function for parsing the body.
// Is any error occurs it will panic.
func ParseBody(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

// ParseBodyValidator is helpers function for parsing the body
// and validation structure
func ParseBodyValidator(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ParseBody(ctx, body); err != nil {
		return err
	}

	validate := validator.New()

	if err := validate.Struct(body); err != nil {
		return fiber.ErrUnprocessableEntity
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

// MapErr is helper function for fiber.Map result
func MapErr(v interface{}) fiber.Map {
	return fiber.Map{"error": v}
}
