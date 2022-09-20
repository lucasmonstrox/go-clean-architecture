package main

import (
	"app/src/ports/fasthttp"
	mongogodriver "app/src/ports/mongo-go-driver"
)

func main() {
	client := mongogodriver.Connect()
	router := fasthttp.MakeRoutes(client)
	// TODO: disable in production
	fasthttp.MakeSwagger(router)
	fasthttp.Start(router)
}
