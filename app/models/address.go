package model

import (
	"childgo/utils/uuidv7"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model

	ID        uuid.UUID `gorm:"type:uuid"`
	City      string    `gorm:"default:Красноярск"  json:"city"`
	Street    string    `gorm:"default:Калинина"  json:"street"`
	House     string    `gorm:"type:string"  json:"house" validate:"required"`
	Apartment string    `gorm:"type:string"  json:"apartment" validate:"required"`

	ChildID uuid.UUID
}

func (a *Address) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, generateError := uuidv7.Generate()

	if generateError != nil {
    return errors.New("can't save invalid data")
	}

  a.ID = *uuid

  return
}