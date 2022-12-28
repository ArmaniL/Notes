package controllers

import (
	"NotesApp/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// to be finished
func GetNoteByIDHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect Query Parameter",
		})
		return
	}
	note, err := db.FindNoteById(id)

	if note.User != c.Request.Header.Get("Email") {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Error Retrieving Note",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Retrieving Note",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"note": note,
	})

}
