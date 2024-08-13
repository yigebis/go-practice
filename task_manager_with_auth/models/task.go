package models

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
