package app

import (
	routes "childgo/app/routes/api/v1"
	"childgo/config"
	storage "childgo/config/database"

	"github.com/gofiber/fiber/v2"
)

func StartupApp(c storage.Connector, o storage.Option) *fiber.App {
	app := fiber.New()

	if err := c.ConnectDB(o.DB); err != nil {
		panic(err)
	}

	if err := c.ConnectCache(o.Cache); err != nil {
		panic(err)
	}

	config.SetupConfigs(app)

	api := app.Group("api/v1")

	routes.AuthRoutes(api)
	routes.ChildRoutes(api)

	return app
}

func StartupTest(c storage.Connector) *fiber.App {
	dbMem := "file::memory:?cache=shared"
	cache := ""

	return StartupApp(c, storage.Option{DB: dbMem, Cache: cache})
}
