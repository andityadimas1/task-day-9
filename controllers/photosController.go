package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"task-day-11/models"

	"github.com/gin-gonic/gin"
)

func (StrDB *StrDB) Photos(c *gin.Context) {
	var (
		albums models.Photos
		result gin.H
	)

	resp, err := http.Get("https://jsonplaceholder.typicode.com/albums/1")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type Photos struct {
		AlbumId      uint   `json:"albumId"`
		ID           uint   `gorm:"primarykey" json:"id"`
		Title        string `json:"title"`
		Url          string `json:"url"`
		ThumbnailUrl string `json:"thumbnailUrl"`
	}

	fmt.Println(body)
	fmt.Println(string(body))
	getData := Photos{}
	if error := json.Unmarshal(body, &getData); error != nil {
		fmt.Println("error", error.Error())
	}
	fmt.Println(getData)

	albums.AlbumID = getData.AlbumId
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
