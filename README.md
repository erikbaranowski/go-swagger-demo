# go-swagger-demo

This repo contains a demo of [go-swagger](https://goswagger.io/)

Most of this code is generated by running `make server`

This code is not fully generated:
- restapi/configure_greeting_server.go is generated but then edited and won't get updated when making the server.
- handlers/get_greeting.go is manually written to implement the functionality for an endpoint.