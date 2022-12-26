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

func FindNoteById(idString string) (*Note, error) {

	note := &Note{}
	err := mgm.Coll(note).FindByID(idString, note)

	return note, err
}

func UpdateNote(note *Note, params UpdateNoteParameters) error {
	data := params.Data
	if params.Parameter == "Header" {
		note.Header = data
	} else if params.Parameter == "Content" {
		note.Content = data
	}

	err := mgm.Coll(note).Update(note)
	return err
}

func ListNotes(email string) ([]Note, error) {
	result := []Note{}
	err := mgm.Coll(&Note{}).SimpleFind(&result, bson.M{"user": email})

	return result, err
}

func GetUser(email, password string) (User, error) {
	user := User{}
	err := mgm.Coll(&User{}).First(bson.M{"user": email}, &user)
	return user, err
}

func DoesUserExist(email string) (bool, error) {
	result := []User{}
	err := mgm.Coll(&User{}).SimpleFind(&result, bson.M{"email": email})
	return len(result) > 0, err
}

func CreateUser(email, password string) error {
	user := NewUser(email, password)
	// Make sure to pass the model by reference (to update the model's "updated_at", "created_at" and "id" fields by mgm).
	err := mgm.Coll(user).Create(user)
	return err
}
