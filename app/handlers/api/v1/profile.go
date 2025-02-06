package handler

import (
	"childgo/utils"
	"github.com/gofiber/fiber/v2"
)

type ProfileResponse struct {
	Email string `json:"email"`
}

func Profile(ctx *fiber.Ctx) error {
	u := utils.GetUser(ctx)

	return ctx.JSON(ProfileResponse{
		Email: u.Email,
	})
}
