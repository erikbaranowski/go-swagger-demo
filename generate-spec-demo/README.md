# generate-spec-demo

THIS IS A WORK IN PROGRESS

This folder contains a demo of [go-swagger](https://goswagger.io/). 
The focus is on generating a swagger spec by running `make spec`.

run command: `go run ./cmd/server --port=23456`

Code from `../generate-code-demo` was used as a starting template for this demo.

## Endpoints

- GET localhost:{PORT}/hello?Name={NAME}

  Takes in an optional name parameter and returns a greeting

- GET localhost:{PORT}/docs

  Swagger UI
