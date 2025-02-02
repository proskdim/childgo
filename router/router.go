package router

import (
	"childgo/config"
	"childgo/handler"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api/v1")

	// healthcheck
	api.Get("/", handler.Ok)

	// auth
	api.Post("/signup", handler.Signup)
	api.Post("/signin", handler.Signin)

	authorizedGroup := api.Group("")

	authorizedGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: config.SecretKey,
		},
		ContextKey: config.ContextKeyUser,
	}))

	// authorized api handlers
	authorizedGroup.Get("/profile", handler.Profile)

	//childs
	authorizedGroup.Get("/childs", handler.Childs)
	authorizedGroup.Get("/child/:id", handler.GetChild)
	authorizedGroup.Post("/child", handler.NewChild)
	authorizedGroup.Delete("/child/:id", handler.DeleteChild)
}
