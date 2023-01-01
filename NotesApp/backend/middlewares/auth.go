package middlewares

import (
	"NotesApp/controllers"
	"NotesApp/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequiredMiddleware(c *gin.Context) {
	var email string
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(403)
		return
	}

	tokenString := authHeader[len("Bearer"):]
	tokenString = strings.TrimSpace(tokenString)
	email, err := controllers.DecodeToken(tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	userExists, err := db.DoesUserExist(email)

	if !userExists {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Request.Header.Add("Email", email)
	c.Next()
}
