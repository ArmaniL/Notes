package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/notes", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})

	})

	r.POST("/note", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)

		if err != nil {
			panic("error reading body")
		}

		body := string(bodyBytes)

		fmt.Println(body)

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})

	})
	port := GoDotEnvVariable("PORT")
	portString := fmt.Sprintf(":%s", port)
	r.Run(portString)
}
