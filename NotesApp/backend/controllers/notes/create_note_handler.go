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
	email := c.Request.Header.Get("Email")
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Reading Request Body ",
		})

	}
	var note model.Note
	err = json.Unmarshal(body, &note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Parsing Body",
		})
		return
	}
	if note.Header == "" || note.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Body",
		})
		return
	}

	err = db.CreateNote(note.Header, note.Content, email)

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
