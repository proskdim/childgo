package database

import (
	"childgo/model"
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

var ErrConnection = errors.New("failed to connect database")
var ErrMigration  = errors.New("failed to create migration")

func ConnectDB() error {
	conn, err := gorm.Open(sqlite.Open("child.db"))

	if err != nil {
		return ErrConnection
	}

	DBConn = conn

	if err := createMigration(); err != nil {
		return ErrMigration
	}

	return nil
}

func createMigration() error {
	return DBConn.AutoMigrate(&model.Child{})
}
