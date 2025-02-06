package repo

import (
	model "childgo/app/models"
	"childgo/config/database"

	"gorm.io/gorm"
)

// FindUser finds a user with given condition
func FindUser(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DBConn.Model(&model.User{}).Take(dest, conds...)
}

// CreateUser create a user entry in the user's table
func CreateUser(user *model.User) *gorm.DB {
	return database.DBConn.Create(user)
}
