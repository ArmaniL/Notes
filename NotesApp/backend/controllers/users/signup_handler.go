package controllers

import (
	"NotesApp/controllers"
	"NotesApp/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	email, password, bodyErr, jsonErr := controllers.ParseUserInfo(c.Request.Body)
	if bodyErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Reading Body",
		})
		return
	}

	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Parsing Payload",
		})
		return
	}

	userDoesExist, serverErr := db.DoesUserExist(email)

	if serverErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in server",
		})
		return
	}

	if userDoesExist {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User Exists already ",
		})
		return

	}

	err := db.CreateUser(email, password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "There was a problem creating a new User ",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesful Sign Up",
	})

}
