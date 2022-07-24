package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotesHandler(c *gin.Context) {

	result, err := ListNotes()

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
func CreateNoteHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Creating Note",
		})

	}
	var note Note
	json.Unmarshal(body, &note)
	fmt.Println(fmt.Sprintf("%s %s %s ", note.Header, note.Content, note.User))
	CreateNote(note.Header, note.Content, note.User)
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
