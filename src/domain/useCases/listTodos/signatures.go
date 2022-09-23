package usecases

import "app/src/domain/entities"

type ListTodosRepository func() []entities.Todo
type ListTodosDeps struct {
	listTodosRepository ListTodosRepository
}
type ListTodosUseCase func() ListTodosOutput
type ListTodosOutput []entities.Todo
