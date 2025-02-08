package migration

import "gorm.io/gorm"

func CreateMigration(db *gorm.DB, models ...interface{}) error {
	return db.AutoMigrate(models...)
}
