package main

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoDB() {
	mongoURL := GoDotEnvVariable("MONGOURL")
	err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI(mongoURL))
	if err != nil {
		panic("There was an error configuring the database")
	}
}

func createNote(header, content, user string) {
	note := NewNote(header, content, user)
	// Make sure to pass the model by reference (to update the model's "updated_at", "created_at" and "id" fields by mgm).
	err := mgm.Coll(note).Create(note)
	if err != nil {
		panic("There was aproblem saving the book")
	}
}

func listNotes() {

}
