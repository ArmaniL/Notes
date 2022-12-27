package model

import (
	"github.com/kamva/mgm/v3"
)

type Note struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	User             string `json:"user" bson:"user"`
	Header           string `json:"header" bson:"header"`
	Content          string `json:"content" bson:"content"`
}

func NewNote(header, content, user string) *Note {
	return &Note{
		User:    user,
		Header:  header,
		Content: content,
	}
}
