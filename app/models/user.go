package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uuid.UUID `gorm:"type:uuid"`
	Email    string    `gorm:"not null" json:"email" validate:"required"`
	Password string    `gorm:"not null" json:"password" validate:"required"`

	Childrens []Child
}
