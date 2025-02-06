package uuidv7

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Generate() (*uuid.UUID, *fiber.Error) {
	uuid, err := uuid.NewV7()

	if err != nil {
		return nil, fiber.ErrBadRequest
	}

	return &uuid, nil
}
