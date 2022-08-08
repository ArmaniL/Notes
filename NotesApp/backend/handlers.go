package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	userDoesExist, serverErr := DoesUserExist(email)

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

	err := CreateUser(email, password)

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

	user, err := GetUser(email, password)
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

func GetNotesHandler(c *gin.Context) {
	var email string
	authHeader := c.Request.Header.Get("Authorization")
	tokenString := authHeader[len("Bearer"):]
	fmt.Println(tokenString)
	email, err := DecodeToken(tokenString)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No user found",
		})
		return
	}

	result, err := ListNotes(email)

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
