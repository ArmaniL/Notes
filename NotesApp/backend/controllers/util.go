package controllers

import (
	"NotesApp/config"
	"NotesApp/model"
	"encoding/json"
	"fmt"
	"io"

	"github.com/golang-jwt/jwt"
)

func ParseUserInfo(Body io.ReadCloser) (string, string, error, error) {
	body, bodyErr := io.ReadAll(Body)
	params := model.User{}
	jsonerr := json.Unmarshal(body, &params)
	email := params.User
	password := params.Password
	return email, password, bodyErr, jsonerr

}

func GenerateToken(user string) (string, error) {
	secret := config.GoDotEnvVariable("SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user,
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(secret))

}

func DecodeToken(token string) (string, error) {

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		secret := config.GoDotEnvVariable("SECRET")
		return []byte(secret), nil
	})

	if err != nil {

		return "", err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims["email"].(string), nil
	}

	return "", err

}
