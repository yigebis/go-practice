package main

import (
	"task_management_api_with_clean_architecture/delivery/router"
	"task_management_api_with_clean_architecture/infrastructure"
	"task_management_api_with_clean_architecture/repository"
	"task_management_api_with_clean_architecture/usecase"

	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var username = "yigerem4"
var password = "yige1234"
var uri = "mongodb+srv://" + username + ":" + password + "@cluster0.isgee.mongodb.net/"

var secretKey = "123456"

func main() {
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

	user_collection := client.Database("test").Collection("users")
	task_collection := client.Database("test").Collection("tasks")

	ur := repository.NewUserRepository(user_collection)
	ps := infrastructure.NewPasswordService()
	ts := infrastructure.NewTokenService(secretKey)

	tr := repository.NewTaskRepository(task_collection)

	uuc := usecase.NewUserUseCase(ur, ps, ts)
	tuc := usecase.NewTaskUseCase(tr)

	router.Setup(uuc, tuc, secretKey)
}
