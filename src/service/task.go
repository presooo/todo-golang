package service

import (
	"context"
	"log"
	"todo/src/db"
	"todo/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func init() {
	var client *mongo.Client = db.DbConnect()
	collection = client.Database(db.DbName).Collection("tasks")
}

func GetTask(task string) bson.M {
	var result bson.M
	id, _ := primitive.ObjectIDFromHex(task)

	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func CreateTask(task models.Tasks) int64 {
	result, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	return result.InsertedID
}

func DeleteTask(task string) int64 {

	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	return result.DeletedCount
}
