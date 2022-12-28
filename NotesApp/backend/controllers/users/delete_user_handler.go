package controllers

import (
	"NotesApp/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect Path Parameter",
		})
		return
	}
	err := db.DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Deleting Note",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})

}
