package entities

import (
	"time"
)

type CreateTodoInput struct {
	Title string
}

type Todo struct {
	Id        string
	Title     string
	Done      bool
	CreatedAt time.Time
	DoneAt    *time.Time
}
