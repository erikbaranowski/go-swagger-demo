package operations

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*
GetGreetingOK returns a greeting

swagger:response getGreetingOK
*/
type GetGreetingOK struct {

	/*contains the actual greeting as plain text
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// WriteResponse to the client
func (o *GetGreetingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
