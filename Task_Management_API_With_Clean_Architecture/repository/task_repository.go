package repository

import (
	"task_management_api_with_clean_architecture/domain"
	"task_management_api_with_clean_architecture/usecase"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	Collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) usecase.ITaskRepository {
	return &TaskRepository{Collection: collection}
}

func (tr *TaskRepository) FetchAllTasks() ([]domain.Task, error) {
	var Tasks []domain.Task

	cursor, err := tr.Collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var task domain.Task

		err = cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		Tasks = append(Tasks, task)
	}

	return Tasks, nil
}
func (tr *TaskRepository) FetchTaskById(id string) (domain.Task, error) {
	var task domain.Task

	filter := bson.D{{Key: "id", Value: id}}

	err := tr.Collection.FindOne(context.TODO(), filter).Decode(&task)

	return task, err
}
func (tr *TaskRepository) InsertTask(newTask domain.Task) error {
	_, err := tr.Collection.InsertOne(context.TODO(), newTask)
	return err
}
func (tr *TaskRepository) UpdateTask(id string, updatedTask domain.Task) error {

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

	_, err := tr.Collection.UpdateOne(context.TODO(), filter, update)

	return err
}
func (tr *TaskRepository) DeleteTask(id string) error {
	filter := bson.D{{Key: "id", Value: id}}

	_, err := tr.Collection.DeleteOne(context.TODO(), filter)

	return err
}
