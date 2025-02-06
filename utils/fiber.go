package utils

import (
	model "childgo/app/models"
	"childgo/config"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
func ParseID(ctx *fiber.Ctx) (int, *fiber.Error) {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return 0, fiber.ErrBadRequest
	}

	return id, nil
}

// GetUser is helper function for getting authenticated user's id
func GetUser(ctx *fiber.Ctx) *model.User {
	id, _ := ctx.Locals(config.ContextJwtUser).(*model.User)
	return id
}
