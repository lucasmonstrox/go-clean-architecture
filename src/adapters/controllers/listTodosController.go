package controllers

import (
	"app/src/adapters"
	usecases "app/src/domain/useCases/listTodos"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

// @Description Get a list of all todos
// @Tags        todos
// @Success     200 {object} schemas.Todo
// @Router      /todos [get]
func SetupListTodosController(listTodosRepository usecases.ListTodosRepository) func(context *fasthttp.RequestCtx) {
	useCase := usecases.SetupListTodosUseCase(listTodosRepository)
	return func(context *fasthttp.RequestCtx) {
		todos := useCase()
		todosToRest := adapters.TodosToRest(todos)
		res, err := json.Marshal(todosToRest)
		hasErrorEncodingJson := err != nil
		if hasErrorEncodingJson {
			fmt.Print("Error encoding json")
		}
		context.SetBody(res)
	}
}
