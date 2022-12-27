package db

import (
	"NotesApp/config"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() {
	mongoURL := config.GoDotEnvVariable("MONGOURL")
	err := mgm.SetDefaultConfig(nil, "test", options.Client().ApplyURI(mongoURL))
	if err != nil {
		panic("There was an error configuring the database")
	}
}
