package fasthttp

import (
	"app/docs"
	"app/src/adapters/controllers"
	"app/src/adapters/repositories"

	"github.com/fasthttp/router"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeRoutes(client *mongo.Client) *router.Router {
	r := router.New()
	listTodosRepository := repositories.SetupListTodosRepository(client)
	r.GET("/todos", controllers.SetupListTodosController(listTodosRepository))
	createTodoRepository := repositories.SetupCreateTodoRepository(client)
	r.POST("/todos", controllers.SetupCreateTodoController(createTodoRepository))
	return r
}

func MakeSwagger(router *router.Router) {
	docs.SwaggerInfo.Title = "Fasthttp Swagger"
	docs.SwaggerInfo.Description = "Fasthttp Swagger"
	docs.SwaggerInfo.Version = "1.0.0-0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = ""
	router.GET("/{filepath:*}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))
}

// @title                     Swagger Example API
// @version                   1.0
// @description               This is a sample server celler server.
// @termsOfService            http://swagger.io/terms/
// @contact.name              API Support
// @contact.url               http://www.swagger.io/support
// @contact.email             support@swagger.io
// @license.name              Apache 2.0
// @license.url               http://www.apache.org/licenses/LICENSE-2.0.html
// @host                      localhost:3000
// @BasePath                  /api/v1
// @securityDefinitions.basic BasicAuth
func Start(router *router.Router) {
	fasthttp.ListenAndServe(":3000", router.Handler)
}
