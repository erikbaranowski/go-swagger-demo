# goa-demo

This folder contains a demo of [goa](https://goa.design/). 

- The `gen` folder is populated via `make design`
- The `cmd` folder and `calc.go` were initialized via `make cmd`

## Server

run server command: `go run ./cmd/calc`

## CLI

run cli http command: `go run ./cmd/calc-cli --url="http://localhost:8000" calc multiply --a 1 --b 2`

run cli grpc command: `go run ./cmd/calc-cli --url="grpc://localhost:8080" calc multiply --message '{"a": 1, "b": 2}'`

## Endpoints

- GET http://localhost:8000/multiply/{A}/{B}

  multiplies A * B

- GET http://localhost:8000/openapi.json

  serves OpenAPI spec for usage in any editor or codegen tool