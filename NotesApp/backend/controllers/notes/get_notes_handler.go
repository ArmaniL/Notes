package controllers

import (
	"NotesApp/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotesHandler(c *gin.Context) {
	email := c.Request.Header.Get("Email")
	result, err := db.ListNotes(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Creating Note",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"notes": result,
	})

}
