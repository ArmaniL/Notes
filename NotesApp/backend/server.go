package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	InitMongoDB()

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/notes", GetNotesHandler)
	r.POST("/notes", CreateNoteHandler)
	r.POST("/login", LoginHandler)
	r.POST("/signup", SignUpHandler)
	port := GoDotEnvVariable("PORT")
	r.Run(port)
}
