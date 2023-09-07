# go-swagger-demo

This repo contains a demo of [go-swagger](https://goswagger.io/). 
Most of this code is generated by running `make server`.

run command: `go run ./cmd/greeting-server-server`

## Not Fully Generated:

- `swagger.yml` the swagger file used for generating all the server code.
- `restapi/configure_greeting_server.go` is generated but then edited and won't get updated when making the server.
- `handlers/get_greeting.go` is manually written to implement the functionality for an endpoint.

## Endpoints

- GET localhost:{PORT}/hello

  Takes in an optional name parameter and returns a greeting

- GET localhost:{PORT}/docs

  Swagger UI
