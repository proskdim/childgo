package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Child struct {
	gorm.Model

	ID       uuid.UUID `gorm:"type:uuid"`
	Name     string    `gorm:"not null"  json:"name" validate:"required"`
	Age      int       `gorm:"not null"  json:"age"  validate:"required,gte=1,lte=18"`
	Birthday time.Time `gorm:"not null"  json:"birthday" validate:"required"`

	UserID uuid.UUID
}
