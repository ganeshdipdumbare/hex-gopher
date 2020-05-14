![Test](https://github.com/ganeshdipdumbare/hex-gopher/workflows/Test/badge.svg) [![codecov](https://codecov.io/gh/ganeshdipdumbare/hex-gopher/branch/master/graph/badge.svg)](https://codecov.io/gh/ganeshdipdumbare/hex-gopher)
# hex-gopher
Simple implementation of hexagonal architecture for Go service. 

## Description

Hexagonal architecture allow us to focus on business logic rather than on tools and external connections. Core business logic is safely developed at domain level.  

-   This is a simple implementation of hexagonal architecture in which core logic is in the app.
-   repo is handling DB related requests only, we can plug and play another DB at anytime if it satisifies interface `GopherDB`
-   api is the entry point to the application, outside system can interact with the microservice using apis and it can be of any form(gRPC,REST).
-   connector is for connectiong the microservice to other microservices(yet to be implemented).
-   The dependency is inward as app is not dependent on any other module which allow us to keep our businees logic intact irrespective of DB,apis etc.

## Usage
-   Clone the repo using  
    `git clone`
-   Start Redis DB server on local machine.
-   Start gRPC server by using following command  
    `go run main.go`
-   You can call gRPC endpoint using bloomRPC.
-   Use proto file which is placed under /api.
-   To run test case-  
    `go test ./...`

## Enhancements
-   Implememtation of REST and GraphQL sever.
-   Add mongoDB.
-   Add sample connection with other microservice.
