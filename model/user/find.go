package user

import (
	"childgo/model"
	"gorm.io/gorm"
)

func Find(db *gorm.DB, u *model.User) (*model.User, error) {
	err := db.Where("email = ? AND password = ?", &u.Email, &u.Password).First(&u).Error
	return u, err
}

func FindByEmail(db *gorm.DB, email string) (*model.User, error) {
	u := &model.User{}
	err := db.Model(u).Where("Email = ?", email).First(u).Error
	return u, err
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

func UpdateChild(db *gorm.DB, source *model.Child, new *model.Child) (*model.Child, error) {
	source.Name = new.Name
	source.Age = new.Age
	source.Birthday = new.Birthday

	err := db.Save(source).Error

	return source, err
}
