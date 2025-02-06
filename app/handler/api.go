package handler

import "github.com/gofiber/fiber/v2"

func Ok(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}
