package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"task_manager_with_database/models"

	// "time"

	// "github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrNoTask     = errors.New("task not found")
	ErrTaskExists = errors.New("task already exists")
)

var username = "yigerem4"
var password = "yige1234"
var uri = "mongodb+srv://" + username + ":" + password + "@cluster0.isgee.mongodb.net/"

func connectToDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to db!")

	return client
}

func getCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("test").Collection("tasks")

	return collection
}

var client = connectToDB()
var collection = getCollection(client)

func GetAllTasks() ([]models.Task, error) {
	var Tasks []models.Task

	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var task models.Task

		err = cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		Tasks = append(Tasks, task)
	}

	return Tasks, nil
}

func GetTaskById(id string) (models.Task, error) {
	var task models.Task
	filter := bson.D{{Key: "id", Value: id}}

	err := collection.FindOne(context.TODO(), filter).Decode(&task)

	if err != nil {
		return models.Task{}, err
	}

	return task, nil

}

func CreateTask(newTask models.Task) error {
	newId := newTask.Id

	//TODO check whether id exists or not
	_, err := GetTaskById(newId)

	if err == nil {
		return ErrTaskExists
	}

	createResult, err := collection.InsertOne(context.TODO(), newTask)

	if err != nil {
		return err
	}

	fmt.Println(createResult.InsertedID)

	return nil
}

func UpdateTask(id string, updatedTask models.Task) error {

	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{}

	if updatedTask.Title != "" {
		update = append(update, bson.E{Key: "$set", Value: bson.D{{Key: "title", Value: updatedTask.Title}}})
	}
	if !updatedTask.DueDate.IsZero() {
		update = append(update, bson.E{Key: "$set", Value: bson.D{{Key: "duedate", Value: updatedTask.DueDate}}})
	}
	if updatedTask.Status != "" {
		update = append(update, bson.E{Key: "$set", Value: bson.D{{Key: "status", Value: updatedTask.Status}}})
	}
	if updatedTask.Description != "" {
		update = append(update, bson.E{Key: "$set", Value: bson.D{{Key: "description", Value: updatedTask.Description}}})
	}

	update_result, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return err
	}

	if update_result.MatchedCount == 0 {
		return ErrNoTask
	}

	return nil
}

func DeleteTask(id string) error {
	filter := bson.D{{Key: "id", Value: id}}

	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return ErrNoTask
	}

	return nil
}
