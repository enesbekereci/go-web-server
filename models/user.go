package models

import (
	"time"
)

type User struct {
	ID            uint       `gorm:"primaryKey"`
	Name          string     //
	Password      string     //
	Description   *string    //
	UserRole      UserRole   //
	KeyValuePairs []KeyValue `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt     time.Time  `gorm:"autoCreatedTime"`
	UpdatedAt     time.Time  `gorm:"autoUpdatedTime"`
}

type UserRole int64

const (
	RoleManager UserRole = 0
	RoleUser    UserRole = 1
)
