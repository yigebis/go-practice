package repository

import (
	"context"
	"task_management_api_with_clean_architecture/domain"
	"task_management_api_with_clean_architecture/usecase"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) usecase.IUserRepository {
	return &UserRepository{Collection: collection}
}

func (ur *UserRepository) GetUser(username string) (domain.User, error) {
	var existingUser domain.User

	filter := bson.D{{Key: "username", Value: username}}

	err := ur.Collection.FindOne(context.TODO(), filter).Decode(&existingUser)

	if err != nil {
		return domain.User{}, err
	}

	return existingUser, nil
}

func (ur *UserRepository) IsDatabaseEmpty() (bool, error) {
	count, err := ur.Collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return false, err
	}

	if count == 0 {
		return true, nil
	}

	return false, nil
}

func (ur *UserRepository) AddUser(user domain.User) error {
	_, err := ur.Collection.InsertOne(context.TODO(), user)

	return err
}

func (ur *UserRepository) UpdateRole(username string) (int64, error) {
	filter := bson.M{"username": username}

	update := bson.M{"$set": bson.M{"role": "admin"}}

	update_result, err := ur.Collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return 0, err
	}

	return update_result.ModifiedCount, nil
}
