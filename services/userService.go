package services

import (
	"go-web-server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func AddUserService() (uint, error) {
	db, err := gorm.Open(sqlite.Open("db/sqlite.db"), &gorm.Config{})
	if err != nil {
		return 0, err
	}
	user := models.User{Name: "user-1", Password: "pass-2"}

	result := db.Create(&user) // pass pointer of data to Create

	return user.ID, result.Error // returns error
}

func DeleteUserService(ID uint) (uint, error) {
	db, err := gorm.Open(sqlite.Open("db/sqlite.db"), &gorm.Config{})
	if err != nil {
		return 0, err
	}
	result := db.Delete(&models.User{}, ID)
	return uint(result.RowsAffected), result.Error // returns error
}

func GetUsersService() ([]models.User, error) {
	db, err := gorm.Open(sqlite.Open("db/sqlite.db"), &gorm.Config{})
	if err != nil {
		return []models.User{}, err
	}
	var users []models.User
	result := db.Find(&users)

	return users, result.Error // returns error
}
