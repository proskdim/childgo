package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Child struct {
	gorm.Model

	ID       uuid.UUID `gorm:"type:uuid"`
	Name     string    `gorm:"not null"  json:"name"`
	Age      int       `gorm:"not null"  json:"age"`
	Birthday time.Time `gorm:"not null"  json:"birthday"`

	UserID uuid.UUID
}
