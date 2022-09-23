package controllers

import (
	"app/src/domain/entities"
	usecases "app/src/domain/useCases/createTodo"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

// @Description Create a new todo
// @Tags        todos
// @Param       payload body     entities.CreateTodoInput true "input"
// @Success     200     {object} entities.Todo
// @Router      /todos [post]
func SetupCreateTodoController(createTodoRepository usecases.CreateTodoRepository) func(context *fasthttp.RequestCtx) {
	useCase := usecases.SetupCreateTodoUseCase(createTodoRepository)
	return func(context *fasthttp.RequestCtx) {
		var input entities.CreateTodoInput
		json.Unmarshal([]byte(string(context.PostBody())), &input)
		fmt.Print(input)
		todo := useCase(input)
		res, errorEncodingJson := json.Marshal(todo)
		hasErrorEncodingJson := errorEncodingJson != nil
		if hasErrorEncodingJson {
			fmt.Print("Error encoding json")
		}
		context.SetBody(res)
	}
}
