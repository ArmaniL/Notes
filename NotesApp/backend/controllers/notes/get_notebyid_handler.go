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

	userEmail := c.Request.Header.Get("Email")
	if note.User != userEmail {
		//must work on this
		userHasAccess, err := db.DoesUserHaveAccessToNote(note, userEmail)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Server Error",
			})
			return

		}

		if !userHasAccess {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Error Retrieving Note",
			})
			return
		}
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
