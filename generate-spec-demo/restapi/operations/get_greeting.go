package operations

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

/*
GetGreeting swagger:route GET /hello getGreeting

GetGreeting get greeting API

	    Responses:
		  200: getGreetingOK
*/
type GetGreeting struct {
	Context *middleware.Context
	Handler GetGreetingHandler
}

func (o *GetGreeting) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetGreetingParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)
}

func GetGreetingsHandler(params GetGreetingParams) middleware.Responder {
	name := swag.StringValue(params.Name)
	if name == "" {
		name = "World"
	}

	greeting := fmt.Sprintf("Hello, %s!", name)
	return NewGetGreetingOK().WithPayload(greeting)
}

// GetGreetingHandlerFunc turns a function with the right signature into a get greeting handler
type GetGreetingHandlerFunc func(GetGreetingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGreetingHandlerFunc) Handle(params GetGreetingParams) middleware.Responder {
	return fn(params)
}

// GetGreetingHandler interface for that can handle valid get greeting params
type GetGreetingHandler interface {
	Handle(GetGreetingParams) middleware.Responder
}

// NewGetGreeting creates a new http.Handler for the get greeting operation
func NewGetGreeting(ctx *middleware.Context, handler GetGreetingHandler) *GetGreeting {
	return &GetGreeting{Context: ctx, Handler: handler}
}
