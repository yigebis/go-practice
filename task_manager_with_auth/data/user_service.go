package data

import (
	"net/http"
	"task_manager_with_auth/models"

	// "net/http"
	"context"
	"time"

	// "fmt"
	"errors"
	// "strings"

	"github.com/dgrijalva/jwt-go"
	// "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var jwt_secret_signer = "123456"

func getUserCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("test").Collection("users")

	return collection
}

var u_client = connectToDB()
var u_collection = getUserCollection(u_client)

func RegisterUser(user models.User) (int, error) {
	var existingUser models.User

	err := u_collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return http.StatusConflict, errors.New("user already exists")
	}

	hashedPassword, hash_err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if hash_err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	user.Role = "user"
	user.Password = string(hashedPassword)

	// check if the user is the first
	count, err := u_collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	if count == 0 {
		user.Role = "admin"
	}

	_, err = u_collection.InsertOne(context.TODO(), user)

	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	return 200, nil
}

func LogUser(username string, password string) (string, int, error) {
	var user models.User

	filter := bson.D{{Key: "username", Value: username}}

	var existingUser models.User

	err := u_collection.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err != nil {
		return "", http.StatusNotFound, errors.New("invalid username or password")
	}

	//check password match
	if bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password)) != nil {
		return "", http.StatusNotFound, errors.New("invalid username or password")
	}

	user.ID = existingUser.ID
	user.Password = existingUser.Password
	user.Role = existingUser.Role
	user.Username = existingUser.Username

	//grant jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	jwtToken, err := token.SignedString([]byte(jwt_secret_signer))

	if err != nil {
		return "", http.StatusInternalServerError, errors.New("internal server error")
	}

	return jwtToken, 200, nil

}

func Promote(username string) (int, error) {
	filter := bson.M{"username": username}

	update := bson.M{"$set": bson.M{"role": "admin"}}

	update_result, err := u_collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	if update_result.MatchedCount == 0 {
		return http.StatusNotFound, errors.New("no such user")
	}
	if update_result.ModifiedCount == 0 {
		return http.StatusConflict, errors.New("user is already an admin")
	}

	return 200, nil

}
