package adapters

import (
	"app/src/domain/entities"
	"app/src/ports/fasthttp/schemas"
)

func TodoToRest(todo entities.Todo) schemas.Todo {
	restTodo := schemas.Todo{
		Id:        todo.Id,
		Title:     todo.Title,
		Done:      todo.Done,
		CreatedAt: todo.CreatedAt,
		DoneAt:    todo.DoneAt,
	}
	return restTodo
}

func TodosToRest(todos []entities.Todo) []schemas.Todo {
	var restTodos []schemas.Todo
	for _, todo := range todos {
		restTodo := TodoToRest(todo)
		restTodos = append(restTodos, restTodo)
	}
	return restTodos
}
