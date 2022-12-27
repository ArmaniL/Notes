package main

import (
	"NotesApp/config"
	"NotesApp/controllers"
	"NotesApp/db"
	"NotesApp/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitMongoDB()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware)
	r.POST("/notes/share", middlewares.AuthRequiredMiddleware, controllers.ShareNoteHandler)
	r.GET("/notes", middlewares.AuthRequiredMiddleware, controllers.GetNotesHandler)
	r.POST("/notes", middlewares.AuthRequiredMiddleware, controllers.CreateNoteHandler)
	r.PUT("/notes", middlewares.AuthRequiredMiddleware, controllers.UpdateNoteHandler)
	r.POST("/login", controllers.LoginHandler)
	r.POST("/signup", controllers.SignUpHandler)
	port := config.GoDotEnvVariable("PORT")
	r.Run(port)
}
