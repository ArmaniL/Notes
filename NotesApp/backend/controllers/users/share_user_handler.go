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

	err = db.ShareNoteWithUser(userEmail, noteID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

}
