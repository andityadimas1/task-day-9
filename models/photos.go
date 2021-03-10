package models

import (
	"time"

	"gorm.io/gorm"
)

// Post struct
type Photos struct {
	AlbumID      uint   `json:"albumId"`
	ID           uint   `gorm:"primarykey" json:"id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
