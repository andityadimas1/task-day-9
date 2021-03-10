package models

import (
	"time"

	"gorm.io/gorm"
)

// Post struct
type Albums struct {
	UserID    uint   `json:"user_id"`
	ID        uint   `gorm:"primarykey" json:"id"`
	Title     string `json:"title"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
