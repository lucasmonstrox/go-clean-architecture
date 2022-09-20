package controllers

import (
	usecases "app/src/domain/useCases/listContents"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

// ShowAccount godoc
// @Summary     List all contents
// @Description Get a list of all contents
// @Tags        contents
// @Accept      json
// @Produce     json
// @Success     200 {object} fasthttp.Content
// @Router      /contents [get]
func SetupListContentsController(listContentsRepository usecases.ListContentsRepository) func(context *fasthttp.RequestCtx) {
	return func(context *fasthttp.RequestCtx) {
		contents := listContentsRepository()
		res, err := json.Marshal(contents)
		hasErrorEncodingJson := err != nil
		if hasErrorEncodingJson {
			fmt.Print("Error encoding json")
		}
		context.SetBody(res)
	}
}
