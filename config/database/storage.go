package storage

import (
	model "childgo/app/models"
	"childgo/utils/migration"
	"context"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Storage = &storage{}

type Connector interface {
	ConnectDB(p string) error
	ConnectCache(u string) error
}

type storage struct {
	DB     *gorm.DB
	Cache  *redis.Client
}

type Option struct {
	DB    string
	Cache string
}

func (s *storage) ConnectDB(p string) error {
	db, err := gorm.Open(sqlite.Open(p))

	if err != nil {
		return err
	}

	s.DB = db

	migration.CreateMigration(db, &model.User{}, &model.Child{})

	return nil
}

func (s *storage) ConnectCache(u string) error {
	r := redis.NewClient(&redis.Options{Addr: u})

	if err := r.Ping(context.Background()).Err(); err != nil {
		return err
	}

	s.Cache = r

	return nil
}
