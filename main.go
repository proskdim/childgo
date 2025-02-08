package main

import (
	"childgo/app"
	"childgo/config"
	s "childgo/config/database"
	"childgo/utils/env"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	dbName := env.Fetch("DB_NAME", "child.db")
	cache  := env.Fetch("CACHE", ":6379")

	a := app.StartupApp(s.Storage, s.Option{DB: dbName, Cache: cache})

	logrus.Error(a.Listen(fmt.Sprintf(":%v", config.Port)))
}
