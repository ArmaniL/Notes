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

// to be finished
func GetNoteByIDHandler(c *gin.Context) {
	id := c.Query("id")
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

func CreateNoteHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Reading Request Body ",
		})

	}
	var note model.Note
	json.Unmarshal(body, &note)
	//fmt.Println(fmt.Sprintf("%s %s %s ", note.Header, note.Content, note.User))
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
