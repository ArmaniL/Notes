package main

import (
	"github.com/kamva/mgm/v3"
	"golang.org/x/crypto/bcrypt"
)

type Note struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	User             string `json:"user" bson:"user"`
	Header           string `json:"header" bson:"header"`
	Content          string `json:"content" bson:"content"`
}

type User struct {
	mgm.DefaultModel `bson:",inline"`
	User             string   `json:"email"  bson:"user"`
	Password         string   `json:"password" bson:"password"`
	SharedWith       []string `bson:"shared_with"`
}

func NewNote(header, content, user string) *Note {
	return &Note{
		User:    user,
		Header:  header,
		Content: content,
	}
}

func NewUser(email, password string) *User {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 5)

	return &User{
		User:       email,
		Password:   string(passwordHash),
		SharedWith: make([]string, 0),
	}
}
