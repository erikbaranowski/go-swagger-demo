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
}

// NewGetGreeting creates a new http.Handler for the get greeting operation
func NewGetGreeting(ctx *middleware.Context) *GetGreeting {
	return &GetGreeting{
		Context: ctx,
	}
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

	res := GetGreetingsHandler(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)
}

func GetGreetingsHandler(params GetGreetingParams) middleware.Responder {
	name := swag.StringValue(params.Name)
	if name == "" {
		name = "World"
	}

	greeting := fmt.Sprintf("Hello, %s!", name)
	return &GetGreetingOK{Payload: greeting}
}
