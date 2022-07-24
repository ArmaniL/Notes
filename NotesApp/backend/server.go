package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	InitMongoDB()

	r := gin.Default()
	r.GET("/notes", GetNotesHandler)
	r.POST("/notes", CreateNoteHandler)
	port := GoDotEnvVariable("PORT")
	r.Run(port)
}
