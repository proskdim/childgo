package storage

import (
	model "childgo/app/models"
	"childgo/utils/migration"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Storage = &storage{}

type Connector interface {
	ConnectDB(p string) error
}

type storage struct {
	DB     *gorm.DB
}

type Option struct {
	DB    string
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