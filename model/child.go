package model

import (
	"time"

	"gorm.io/gorm"
)

type Child struct {
	gorm.Model

	Name     string    `gorm:"not null" json:"name"`
	Age      int       `gorm:"not null" json:"age"`
	Birthday time.Time `gorm:"not null" json:"birthday"`
}
