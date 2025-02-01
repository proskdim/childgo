package router

import (
	"childgo/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api/v1")

	api.Get("/", handler.Ok)
}
