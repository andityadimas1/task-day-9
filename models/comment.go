package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	PostID    uint   `json:"postId"`
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Body      string `json:"body"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
