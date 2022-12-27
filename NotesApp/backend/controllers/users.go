package controllers

import (
	"NotesApp/db"
	"NotesApp/model"
	"encoding/json"
	"io"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUpHandler(c *gin.Context) {
	email, password, bodyErr, jsonErr := ParseUserInfo(c.Request.Body)
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

func LoginHandler(c *gin.Context) {

	email, password, bodyErr, jsonErr := ParseUserInfo(c.Request.Body)
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
	token, _ := GenerateToken(email)

	c.SetCookie("token", token, 8640, "/", "localhost", true, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Succesful Login",
		"token":   token,
	})

}

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
