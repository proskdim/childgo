package user

import (
	"childgo/model"
	"errors"

	"gorm.io/gorm"
)

var ErrEmailNotFound = errors.New("record by email not found")

func FindByEmail(db *gorm.DB, email string) (*model.User, error) {
	var user model.User
	res := db.Find(&user, &model.User{Email: email})

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, ErrEmailNotFound
	}

	return &user, nil
}
