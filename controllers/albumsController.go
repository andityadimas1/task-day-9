package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
	"task-day-11/models"

	"github.com/gin-gonic/gin"
)

func (StrDB *StrDB) Albums(c *gin.Context) {
	var (
		albums models.Albums
		result gin.H
	)

	resp, err := http.Get("https://jsonplaceholder.typicode.com/albums/1")
	if err != nil {
		log.Fatal(err)
	}
	title, err := ioutil.ReadAll(resp.Title)
	if err != nil {
		log.Fatalln(err)
	}

	type Albums struct {
		UserId uint   `json:"userId"`
		ID     uint   `gorm:"primarykey" json:"id"`
		Title  string `json:"title"`
	}

	fmt.Println(title)
	fmt.Println(string(title))
	getData := Albums{}
	if error := json.Unmarshal(title, &getData); error != nil {
		fmt.Println("error", error.Error())
	}
	fmt.Println(getData)

	albums.UserId = getData.UserId
	albums.ID = getData.ID
	albums.Title = getData.Title
	StrDB.DB.Create(&albums)

	result = gin.H{
		"status":  "success",
		"message": "Sucessfully Get and Insert Data!",
		"data":    albums,
	}
	c.JSON(http.StatusOK, result)
}
