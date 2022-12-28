package controllers

import (
	"NotesApp/controllers"
	"NotesApp/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c *gin.Context) {

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

	user, err := db.GetUser(email, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Finding User",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "incorrect password",
		})
		return
	}
	token, _ := controllers.GenerateToken(email)

	c.SetCookie("token", token, 8640, "/", "localhost", true, true)
	c.JSON(http.StatusOK, gin.H{
		"message":      "Succesful Login",
		"token":        token,
		"shared_notes": user.SharedNotes,
	})

}
