package handler

import (
	"childgo/app/types"
	"childgo/utils"

	"github.com/gofiber/fiber/v2"
)

func Profile(ctx *fiber.Ctx) error {
	u := utils.GetUser(ctx)

	return ctx.JSON(&types.ProfileResponse{
		Email: u.Email,
	})
}
