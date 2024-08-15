package domain

import (
	"time"
)

type Task struct {
	Id          string    `json:"id" bson:"id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	DueDate     time.Time `json:"duedate" bson:"duedate"`
	Status      string    `json:"status" bson:"status"`
}

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}