package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Auth(c *gin.Context) {
	if err := godotenv.Load("env"); err != nil {
		log.Println("Not found", err.Error())
		result := gin.H{
			"message": "something wrong",
			"token":   nil,
		}
		c.JSON(http.StatusUnauthorized, result)
	}
	secret := os.Getenv("SECRET_TOKEN")
	tokenStringHeader := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenStringHeader, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Method not found or not Hs256", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		log.Println(err.Error())
		result := gin.H{
			"message": "Cannot Access that Operation",
			"eror":    err.Error(),
		}

		c.JSON(http.StatusUnauthorized, result)
		c.Abort()

	} else if token != nil {
		fmt.Println("token verified")
	}
}
