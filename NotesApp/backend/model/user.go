package model

import (
	"github.com/kamva/mgm/v3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	User             string   `json:"email"  bson:"user"`
	Password         string   `json:"password" bson:"password"`
	SharedNotes      []string `bson:"shared_notes"`
}

func NewUser(email, password string) *User {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 5)

	return &User{
		User:        email,
		Password:    string(passwordHash),
		SharedNotes: make([]string, 0),
	}

}
