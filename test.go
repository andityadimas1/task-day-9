package main

import (
	"log"
	"task-day-11/config"
	"task-day-11/controllers"
	"task-day-11/models"

	"github.com/gin-gonic/gin"
)

func main() {
	dbPG := config.Connection()
	strDB := controllers.StrDB{DB: dbPG}
	models.Migrations(dbPG)
	routing := gin.Default()

	//routes
	routing.GET("/comment", strDB.Comment)
	routing.GET("/photos:ID", strDB.Photos)
	routing.GET("/albums", strDB.Albums)
	log.Println("Server up and run on Port 8080")
	routing.Run()
}
