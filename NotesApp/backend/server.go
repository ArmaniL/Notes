package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	InitMongoDB()

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/notes", AuthRequiredMiddleware, GetNotesHandler)
	r.POST("/notes", AuthRequiredMiddleware, CreateNoteHandler)
	r.PUT("/notes", AuthRequiredMiddleware, UpdateNoteHandler)
	r.POST("/login", LoginHandler)
	r.POST("/signup", SignUpHandler)
	port := GoDotEnvVariable("PORT")
	r.Run(port)
}
