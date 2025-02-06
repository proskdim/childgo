package repo

import (
	model "childgo/app/models"
	"childgo/config/database"

	"gorm.io/gorm"
)

// CreateChild create a child entry in the child's table
func CreateChild(child *model.Child) *gorm.DB {
	return database.DBConn.Create(child)
}

// FindChild finds a child with given condition
func FindChild(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DBConn.Model(&model.Child{}).Take(dest, conds...)
}

// FindChildByUser finds a child with given child and user identifier
func FindChildByUser(dest interface{}, childIden interface{}, userIden interface{}) *gorm.DB {
	return FindChild(dest, "id = ? AND user_id = ?", childIden, userIden)
}

// FindChildrensByUser finds the childrens with user's identifier given
func FindChildrensByUser(dest interface{}, userIden interface{}) *gorm.DB {
	return database.DBConn.Model(&model.Child{}).Find(dest, "user_id = ?", userIden)
}

// DeleteChild deletes a child from childrens' table with the given child and user identifier
func DeleteChild(childIden interface{}, userIden interface{}) *gorm.DB {
	return database.DBConn.Unscoped().Delete(&model.Child{}, "id = ? AND user_id = ?", childIden, userIden)
}

// UpdateChild allows to update the child with the given childID and userID
func UpdateChild(childIden interface{}, userIden interface{}, data interface{}) *gorm.DB {
	return database.DBConn.Model(&model.Child{}).Where("id = ? AND user_id = ?", childIden, userIden).Updates(data)
}
