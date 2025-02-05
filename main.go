package main

import (
	"childgo/config"
	"childgo/database"
	"childgo/router"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
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

	logrus.Error(webApp.Listen(fmt.Sprintf(":%v", config.PORT)))
}
