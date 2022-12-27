package middlewares

import (
	"NotesApp/controllers"
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
	//fmt.Printf("token:%s\n", tokenString)
	email, err := controllers.DecodeToken(tokenString)
	if err != nil {
		c.AbortWithStatus(403)
		return
	}
	c.Request.Header.Add("Email", email)
	c.Next()
}
