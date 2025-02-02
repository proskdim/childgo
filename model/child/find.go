package child

import (
	"childgo/model"
	"errors"

	"gorm.io/gorm"
)

var ErrRecordNotFound = errors.New("not found record by id")

func FindById(db *gorm.DB, id int) (*model.Child, error) {
	var child model.Child

	res := db.First(&child, id)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}

	return &child, nil
}