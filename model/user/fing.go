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

func FindAllChilds(db *gorm.DB, user *model.User) (*[]model.Child, error) {
	err := db.Preload("Childrens").First(&user, &user.ID).Error
	return &user.Childrens, err
}

func AddChild(db *gorm.DB, user *model.User, child *model.Child) (*model.Child, error) {
	err := db.Model(&user).Association("Childrens").Append(child)
	return child, err
}
