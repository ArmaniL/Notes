package db

import (
	"NotesApp/model"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func FindUserById(idString string) (*model.User, error) {

	user := &model.User{}
	err := mgm.Coll(user).FindByID(idString, user)

	return user, err
}

func GetUser(email, password string) (model.User, error) {
	user := model.User{}
	err := mgm.Coll(&model.User{}).First(bson.M{"user": email}, &user)
	return user, err
}

func ShareNoteWithUser(email string, noteId string) error {

	result := []model.User{}
	err := mgm.Coll(&model.User{}).SimpleFind(&result, bson.M{"user": email})

	if err != nil {
		return err
	}

	user := result[0]

	user.SharedNotes = append(user.SharedNotes, noteId)

	err = mgm.Coll(&user).Update(&user)
	return err
}

func DoesUserExist(email string) (bool, error) {
	result := []model.User{}
	err := mgm.Coll(&model.User{}).SimpleFind(&result, bson.M{"user": email})
	return len(result) > 0, err
}

func CreateUser(email, password string) error {
	user := model.NewUser(email, password)
	// Make sure to pass the model by reference (to update the model's "updated_at", "created_at" and "id" fields by mgm).
	err := mgm.Coll(user).Create(user)
	return err
}

func GetUserIdFromEmail(email string) (string, error) {
	result := []model.User{}
	err := mgm.Coll(&model.User{}).SimpleFind(&result, bson.M{"user": email})
	return result[0].ID.String(), err
}

func DeleteUser(id string) error {
	user, err := FindUserById(id)
	if err != nil {
		return err
	}
	err = mgm.Coll(user).Delete(user)
	return err
}
