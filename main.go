package main

import (
	"childgo/database"
	"childgo/router"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	webApp := fiber.New()

	if err := database.ConnectDB(); err != nil {
		panic(err)
	}

	router.SetupRoutes(webApp)

	logrus.Error(webApp.Listen(":8080"))
}
