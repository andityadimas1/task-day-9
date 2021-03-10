package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"task-day-11/models"

	"github.com/gin-gonic/gin"
)

func (StrDB *StrDB) Comment(c *gin.Context) {
	var (
		comment models.Comment
		result  gin.H
	)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/comments/10")

	if err != nil {
		fmt.Println(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Comment struct
	type Comment struct {
		PostID uint   `json:"post_id"`
		ID     uint   `gorm:"primarykey" json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Body   string `json:"body"`
	}

	fmt.Println(body)
	fmt.Println(string(body))
	data := Comment{}

	if error := json.Unmarshal(body, &data); error != nil {
		fmt.Println("Error ", err.Error())
	}
	fmt.Println(data)

	comment.PostID = data.PostID
	comment.ID = data.ID
	comment.Name = data.Name
	comment.Email = data.Email
	comment.Body = data.Body
	StrDB.DB.Create(&comment)

	result = gin.H{
		"status": "success",
		"data":   comment,
	}

	c.JSON(http.StatusOK, result)
}
