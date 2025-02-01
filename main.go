package main

import (
	"childgo/router"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	webApp := fiber.New()

	router.SetupRoutes(webApp)

	logrus.Error(webApp.Listen(":8080"))
}
