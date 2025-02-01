package database

import (
	"childgo/model"
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB
var RDconn *redis.Client

const (
	tableName = "child.db"
	cacheAddr = "localhost:6379"
)

var (
	ErrDBConnection = errors.New("failed to connect database")
	ErrRDConnection = errors.New("failed to connect cache")
	ErrMigration    = errors.New("failed to create migration")
)

func ConnectDB() (err error) {
	conn, err := gorm.Open(sqlite.Open(tableName))

	if err != nil {
		return ErrDBConnection
	}

	DBConn = conn

	if err = createMigration(); err != nil {
		return ErrMigration
	}

	return
}

func ConnectCache() (err error) {
	RDconn = redis.NewClient(&redis.Options{Addr: cacheAddr})

	if err = RDconn.Ping(context.Background()).Err(); err != nil {
		return ErrRDConnection
	}

	return
}

func createMigration() error {
	return DBConn.AutoMigrate(
		&model.Child{},
		&model.User{},
	)
}
