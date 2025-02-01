package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupConfigs(app *fiber.App) {
	app.Use(logger.New())
}