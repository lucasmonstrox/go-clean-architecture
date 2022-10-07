package schemas

import (
	"time"
)

type CreateTodoInput struct {
	Title string `json:"title"`
}

type Todo struct {
	Id        string     `json:"id"`
	Title     string     `json:"title"`
	Done      bool       `json:"done"`
	CreatedAt time.Time  `json:"createdAt"`
	DoneAt    *time.Time `json:"doneAt"`
}
