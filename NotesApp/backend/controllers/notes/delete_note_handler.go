package controllers

import (
	"NotesApp/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteNoteHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect Path Parameter",
		})
		return
	}

	note, err := db.FindNoteById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Retrieving Note",
		})
		return
	}

	if note.User != c.Request.Header.Get("Email") {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Error Retrieving Note",
		})
		return
	}

	err = db.DeleteNote(note)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Deleting Note",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Note deleted",
	})

}
