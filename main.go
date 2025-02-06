package main

import (
	routes "childgo/app/routes/api/v1"
	"childgo/config"
	"childgo/config/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()

	if err := database.ConnectDB(); err != nil {
		panic(err)
	}

	if err := database.ConnectCache(); err != nil {
		panic(err)
	}

	config.SetupConfigs(app)

	api := app.Group("api/v1")

	routes.AuthRoutes(api)
	routes.ChildRoutes(api)

	logrus.Error(app.Listen(fmt.Sprintf(":%v", config.Port)))
}
