package routes

import (
	handler "childgo/app/handlers/api/v1"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router) {
	app.Get("/", handler.Ok)

	app.Post("/signup", handler.Signup)
	app.Post("/signin", handler.Signin)
}
