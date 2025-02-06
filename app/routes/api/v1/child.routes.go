package routes

import (
	handler "childgo/app/handlers/api/v1"
	"childgo/config"
	"childgo/utils/middleware"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var jwtConfig = jwtware.Config{
	SigningKey: jwtware.SigningKey{ Key: []byte(config.SecretKey)},
	ContextKey: config.ContextKeyUser,
}

func ChildRoutes(app fiber.Router) {
	authorizedGroup := app.Group("")

	authorizedGroup.Use(jwtware.New(jwtConfig))

	// authorized api handlers
	authorizedGroup.Get("/profile", handler.Profile)
	authorizedGroup.Use(middleware.JwtUserMiddleware)

	authorizedGroup.Get("/childs", handler.Childs)
	authorizedGroup.Get("/child/:id", handler.GetChild)
	authorizedGroup.Post("/child", handler.NewChild)
	authorizedGroup.Delete("/child/:id", handler.DeleteChild)
	authorizedGroup.Patch("/child/:id", handler.UpdateChild)
}