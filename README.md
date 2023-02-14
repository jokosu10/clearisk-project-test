## Overview

## Tutorial

We provide another documentation using swagger to make it easy to understand the basic of API.

How to run this project :

1. git clone `https://github.com/jokosu10/clearisk-project-test.git`
2. cd `clearisk-project-test` and run `go install`
3. run command `go run main.go` to running this server like this `localhost:8000`
4. run command `go test ./controller/ -cover` for running unit testing

## HTTP requests

There are 4 basic HTTP requests that you can use in this API:

* `POST` Create a resource
* `PUT` Update a resource
* `GET` Get a resource or list of resources
* `DELETE` Delete a resource

## HTTP Responses

Each response will include a code(repsonse code),message,status and data object that can be single object or array depending on the query.

## HTTP Response Codes

Each response will be returned with one of the following HTTP status codes:

* `200` `OK` The request was successful
* `400` `Bad Request` There was a problem with the request (security, malformed, data validation, etc.)
* `404` `Not found` An attempt was made to access a resource that does not exist in the API
* `500` `Server Error` An error on the server occurred
