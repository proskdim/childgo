package repo

import (
	model "childgo/app/models"
	s "childgo/config/database"

	"gorm.io/gorm"
)

// FindUser finds a user with given condition
func FindUser(dest interface{}, conds ...interface{}) *gorm.DB {
	return s.Storage.DB.Model(&model.User{}).Take(dest, conds...)
}

// CreateUser create a user entry in the user's table
func CreateUser(user *model.User) *gorm.DB {
	return s.Storage.DB.Create(user)
}

// DeleteUsers delete all records from user's table
func DeleteUsers(db *gorm.DB) {
	db.Exec("DELETE fROM users")
}
