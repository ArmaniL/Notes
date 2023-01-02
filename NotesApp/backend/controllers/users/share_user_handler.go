package controllers

import (
	"NotesApp/db"
	"NotesApp/model"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShareNoteHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Reading Request Body ",
		})

	}
	var payload model.ShareNotePayload
	json.Unmarshal(body, &payload)

	userEmail := payload.Email
	noteID := payload.NoteID

	userExists, err := db.DoesUserExist(userEmail)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if !userExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not Real User",
		})
		return

	}

	err = db.ShareNoteWithUser(userEmail, noteID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})

}
