package handler

import "github.com/gofiber/fiber/v2"

// healthcheck function
func Ok(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}
