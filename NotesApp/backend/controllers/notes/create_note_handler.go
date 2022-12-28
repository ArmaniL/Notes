package controllers

import (
	"NotesApp/db"
	"NotesApp/model"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNoteHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Reading Request Body ",
		})

	}
	var note model.Note
	json.Unmarshal(body, &note)
	db.CreateNote(note.Header, note.Content, note.User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Creating Note",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Note Created",
	})

}
