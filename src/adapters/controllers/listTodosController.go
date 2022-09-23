package controllers

import (
	usecases "app/src/domain/useCases/listTodos"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

// @Description Get a list of all todos
// @Tags        todos
// @Success     200 {object} entities.Todo
// @Router      /todos [get]
func SetupListTodosController(listTodosRepository usecases.ListTodosRepository) func(context *fasthttp.RequestCtx) {
	useCase := usecases.SetupListTodosUseCase(listTodosRepository)
	return func(context *fasthttp.RequestCtx) {
		todos := useCase()
		res, err := json.Marshal(todos)
		hasErrorEncodingJson := err != nil
		if hasErrorEncodingJson {
			fmt.Print("Error encoding json")
		}
		context.SetBody(res)
	}
}
