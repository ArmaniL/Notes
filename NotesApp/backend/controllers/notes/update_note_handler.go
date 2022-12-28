package controllers

import (
	"NotesApp/db"
	"NotesApp/model"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateNoteHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Reading Request Body ",
		})

	}
	var params model.UpdateNotePayload
	json.Unmarshal(body, &params)

	note, err := db.FindNoteById(params.NoteID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Updating Note",
		})
		return
	}

	if note.User != c.Request.Header.Get("Email") {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Error Retrieving Note",
		})
		return
	}

	err = db.UpdateNote(note, params)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Updating Note",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Note Updated Succesfully",
	})

}
