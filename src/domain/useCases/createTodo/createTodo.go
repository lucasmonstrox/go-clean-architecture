package usecases

import "app/src/domain/entities"

func SetupCreateTodoUseCase(createTodoRepository CreateTodoRepository) CreateTodoUseCase {
	return func(input entities.CreateTodoInput) entities.Todo {
		return createTodoRepository(input)
	}
}
