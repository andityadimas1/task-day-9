package main

import (
	"task-day-11/config"
	"task-day-11/controllers"
	"task-day-11/models"

	"github.com/gin-gionic/gin"
)

func main() {
	dbPG := config.Connection()
	strDB := controllers.StrDB{DB: dbPG}

	request := gin.Default()

	models.Migrations(dbPG)

	routing := gin.Default()

	//routes
	routing.GET("/comment", strDB.Comment)
	routing.GET("/albums", strDB.Albums)

	request.Run()
}