package controllers

import (
	"app/src/adapters"
	"app/src/domain/entities"
	usecases "app/src/domain/useCases/createTodo"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

// @Description Create a new todo
// @Tags        todos
// @Param       payload body     schemas.CreateTodoInput true "input"
// @Success     200     {object} schemas.Todo
// @Router      /todos [post]
func SetupCreateTodoController(createTodoRepository usecases.CreateTodoRepository) func(context *fasthttp.RequestCtx) {
	useCase := usecases.SetupCreateTodoUseCase(createTodoRepository)
	return func(context *fasthttp.RequestCtx) {
		var input entities.CreateTodoInput
		json.Unmarshal([]byte(string(context.PostBody())), &input)
		todo := useCase(input)
		todoToRest := adapters.TodoToRest(todo)
		res, errorEncodingJson := json.Marshal(todoToRest)
		hasErrorEncodingJson := errorEncodingJson != nil
		if hasErrorEncodingJson {
			fmt.Print("Error encoding json")
		}
		context.SetBody(res)
	}
}
