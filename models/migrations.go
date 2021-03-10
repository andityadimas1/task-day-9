package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {

	if check := db.Migrator().HasTable(&Photos{}); !check {
		db.Migrator().CreateTable(&Photos{})
		fmt.Println("Post Table has been created!")
	}

	if check := db.Migrator().HasTable(&Comment{}); !check {
		db.Migrator().CreateTable(&Comment{})
		fmt.Println("Comment Table has been created!")
	}

	if check := db.Migrator().HasTable(&Albums{}); !check {
		db.Migrator().CreateTable(&Albums{})
		fmt.Println("Todos Table has been created!")
	}

}
