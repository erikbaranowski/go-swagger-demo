package handlers

import (
	"fmt"
	"go-swagger-demo/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

var GetGreetingHandler = func(params operations.GetGreetingParams) middleware.Responder {
	return getGreetingsHandler(params)
}

func getGreetingsHandler(params operations.GetGreetingParams) middleware.Responder {
	name := swag.StringValue(params.Name)
	if name == "" {
		name = "World"
	}

	greeting := fmt.Sprintf("Hello, %s!", name)
	return operations.NewGetGreetingOK().WithPayload(greeting)
}
