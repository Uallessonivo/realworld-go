package config

import (
	"Github.com/Uallessonivo/RealWorld/internal/core/domain/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("realworld.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if er := db.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Comment{},
		&models.Article{}); er != nil {
		panic("failed to migrate schemas")
	}
}
