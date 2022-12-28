package db

import (
	"NotesApp/model"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateNote(header, content, user string) error {
	note := model.NewNote(header, content, user)
	// Make sure to pass the model by reference (to update the model's "updated_at", "created_at" and "id" fields by mgm).
	err := mgm.Coll(note).Create(note)
	return err
}

func FindNoteById(idString string) (*model.Note, error) {

	note := &model.Note{}
	err := mgm.Coll(note).FindByID(idString, note)

	return note, err
}

func UpdateNote(note *model.Note, params model.UpdateNotePayload) error {
	data := params.Data
	if params.Parameter == "Header" {
		note.Header = data
	} else if params.Parameter == "Content" {
		note.Content = data
	}

	err := mgm.Coll(note).Update(note)
	return err
}

func ListNotes(email string) ([]model.Note, error) {
	result := []model.Note{}
	err := mgm.Coll(&model.Note{}).SimpleFind(&result, bson.M{"user": email})

	return result, err
}

func DeleteNote(note *model.Note) error {
	err := mgm.Coll(note).Delete(note)
	return err
}
