package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTodoInput struct {
	Title string `json:"title"`
}

type Todo struct {
	Id        primitive.ObjectID `json:"id"`
	Title     string             `json:"title"`
	Done      bool               `json:"done"`
	CreatedAt time.Time          `json:"createdAt"`
	DoneAt    *time.Time         `json:"doneAt"`
}
