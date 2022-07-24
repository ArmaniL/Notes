package main

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() {
	mongoURL := GoDotEnvVariable("MONGOURL")
	err := mgm.SetDefaultConfig(nil, "test", options.Client().ApplyURI(mongoURL))
	if err != nil {
		panic("There was an error configuring the database")
	}
}

func CreateNote(header, content, user string) error {
	note := NewNote(header, content, user)
	// Make sure to pass the model by reference (to update the model's "updated_at", "created_at" and "id" fields by mgm).
	err := mgm.Coll(note).Create(note)
	return err
}

func UpdateNote(header, content, user string) error {
	note := NewNote(header, content, user)
	err := mgm.Coll(note).Update(note)
	return err
}

func ListNotes() ([]Note, error) {
	result := []Note{}
	err := mgm.Coll(&Note{}).SimpleFind(&result, bson.M{})

	return result, err
}
