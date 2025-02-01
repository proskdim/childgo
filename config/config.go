package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	SecretKey      = []byte("qwerty123456")
	ContextKeyUser = "user"
)

func SetupConfigs(app *fiber.App) {
	app.Use(logger.New())
}