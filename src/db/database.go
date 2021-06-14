package db

import (
	"context"
	"log"
	"todo/src/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = "mongodb://" + utils.EnvVariable("DATABASE_HOST") + ":27017"

// Database Name
const DbName = "todo"

var Client *mongo.Client

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to the Db
	db, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	pingResult := db.Ping(context.TODO(), nil)
	if pingResult != nil {
		log.Fatal(pingResult)
	}

	Client = db
}

func DbConnect() *mongo.Client {
	return Client
}
