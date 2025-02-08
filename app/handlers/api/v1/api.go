package handler

import (
	"childgo/app/types"

	"github.com/gofiber/fiber/v2"
)

// healthcheck function
func Health(ctx *fiber.Ctx) error {
	return ctx.JSON(&types.HealthResp{
		Status: "ok",
	})
}
