package db

import (
	"go-web-server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Seed() {
	db, _ := gorm.Open(sqlite.Open("db/sqlite.db"), &gorm.Config{})
	db.AutoMigrate(&models.User{})
}
