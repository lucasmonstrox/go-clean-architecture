package usecases

import "app/src/domain/entities"

type CreateTodoRepository func(input entities.CreateTodoInput) entities.Todo
type CreateTodoDeps struct {
	createTodoRepository CreateTodoRepository
}
type CreateTodoUseCase func(input entities.CreateTodoInput) entities.Todo
