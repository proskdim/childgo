package user

import (
	"childgo/model"

	"gorm.io/gorm"
)

func FindByEmail(db *gorm.DB, email string) (user *model.User, err error) {
	err = db.Find(&user, &model.User{Email: email}).Error
	return user, err
}

func FindAllChilds(db *gorm.DB, user *model.User) (*[]model.Child, error) {
	err := db.Preload("Childrens").First(&user, &user.ID).Error
	return &user.Childrens, err
}

func AddChild(db *gorm.DB, user *model.User, child *model.Child) (*model.Child, error) {
	err := db.Model(&user).Association("Childrens").Append(child)
	return child, err
}

func FindChild(db *gorm.DB, user *model.User, childId int) (*model.Child, error) {
	child := &model.Child{}
	err := db.Model(child).Where("user_id = ?", user.ID).First(child, childId).Error
	return child, err
}

func DeleteChild(db *gorm.DB, child *model.Child) (*model.Child, error) {
	err := db.Model(model.Child{}).Delete(child).Error
	return child, err
}
