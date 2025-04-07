package config

import (
	"Cinema_API/entity"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Movie{},
	)

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
