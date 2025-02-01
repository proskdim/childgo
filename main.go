package main

import (
	"childgo/config"
	"childgo/database"
	"childgo/router"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

const (
	appAddress = "localhost:8080"
)

func main() {
	webApp := fiber.New()

	if err := database.ConnectDB(); err != nil {
		panic(err)
	}

	if err := database.ConnectCache(); err != nil {
		panic(err)
	}

	config.SetupConfigs(webApp)
	router.SetupRoutes(webApp)

	logrus.Error(webApp.Listen(appAddress))
}
