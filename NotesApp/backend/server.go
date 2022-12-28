package main

import (
	"NotesApp/config"
	notes "NotesApp/controllers/notes"
	users "NotesApp/controllers/users"
	"NotesApp/db"
	"NotesApp/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitMongoDB()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware)
	r.POST("/notes/share", middlewares.AuthRequiredMiddleware, users.ShareNoteHandler)
	r.GET("/notes", middlewares.AuthRequiredMiddleware, notes.GetNotesHandler)
	r.GET("/notes/:id", middlewares.AuthRequiredMiddleware, notes.GetNoteByIDHandler)
	r.POST("/notes", middlewares.AuthRequiredMiddleware, notes.CreateNoteHandler)
	r.PUT("/notes", middlewares.AuthRequiredMiddleware, notes.UpdateNoteHandler)
	r.DELETE("/notes/:id", middlewares.AuthRequiredMiddleware, notes.DeleteNoteHandler)
	r.POST("/login", users.LoginHandler)
	r.POST("/signup", users.SignUpHandler)
	port := config.GoDotEnvVariable("PORT")
	r.Run(port)
}
