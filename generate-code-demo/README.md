# generate-code-demo

This folder contains a demo of [go-swagger](https://goswagger.io/). 
The focus is on generating code from a swagger.yml by running `make server`.

run command: `go run ./cmd/greeting-server-server`

## Editable

`restapi/configure_greeting_server.go` is generated but then edited and won't get updated when making the server.

## Created

`handlers/get_greeting.go` is manually written for the implementation of the endpoint and is wired in to `restapi/configure_greeting_server.go`.

## Endpoints

- GET localhost:{PORT}/hello?Name={NAME}

  Takes in an optional name parameter and returns a greeting

- GET localhost:{PORT}/docs

  Swagger UI