package models

import "time"

type KeyValue struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      //
	Key       string    // A regular string field
	Value     *string   // A pointer to a string, allowing for null values
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
}
