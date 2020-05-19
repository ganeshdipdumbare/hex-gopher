![Test](https://github.com/ganeshdipdumbare/hex-gopher/workflows/Test/badge.svg) [![codecov](https://codecov.io/gh/ganeshdipdumbare/hex-gopher/branch/master/graph/badge.svg)](https://codecov.io/gh/ganeshdipdumbare/hex-gopher)
# hex-gopher
Simple implementation of hexagonal architecture for Go service. 

## Contents
-   [Description](#Description)
-   [Documentation](#Documentation)
-   [Setup](#Setup)
-   [Testing](#Testing)
-   [Configuration](#Configuration)
-   [Roadmap](#Roadmap)
-   [Owner](#Owner)



## Description

Hexagonal architecture allow us to focus on business logic rather than on tools and external connections. Core business logic is safely developed at domain level. All the other services, tools, interfaces and repositories are connected to domain layer by application layer. The dependecies can be visualized as follows-  

```tools/libraries/repos  -depends on-> application layer -dependes on-> domain(business logic)```  

Sample hex architecture layout is as follows - 
![Image of architecture](https://github.com/ganeshdipdumbare/hex-gopher/hexarchitecture.png)

So dependecy is always inwards towards domain layer. Domain layer/business logic is independednt of all the other layers and can be managed easily. This allow us to change all the other layers(e.g. DB or APIs) without having any changes in core business logic


-   This is a simple implementation of hexagonal architecture in which core logic is in the app.
-   repo is handling DB related requests only, we can plug and play another DB at anytime if it satisifies interface `GopherDB`
-   api is the entry point to the application, outside system can interact with the microservice using apis and it can be of any form(gRPC,REST).
-   connector is for connectiong the microservice to other microservices(yet to be implemented).
-   The dependency is inward as app is not dependent on any other module which allow us to keep our businees logic intact irrespective of DB,apis etc.

## Documentation
This service provides gRPC implementation of GopherService, it provides us following functionality-  

- SaveGopher - Accepts and save Gopher to DB which is setup for the miscroservice. Can be chosen from Redis or MongoDB in the grpc service. Gopher is having following fields-
    ```golang
    type Gopher struct {
        Id string 
        Name string
    } 
    ```
- GetGopher - Get gopher from DB for given ID.

## Requirements
-   RedisDB server
-   MongoDB Server
-   Go compiler

## Setup
-   Clone the repo
-   Start the DB on local machine
-   Build the code using ```go build```
-   Run the program using ```./hex-gopher```
-   Use bloomRPC with proto file in api/grpcapi/grpcapi.proto file to test grpc server  

## Testing
-   Start the DB server on the local machine
-   Run command ```go test ./...```

## Configuration
-   Start MongoDB or RedisDB server as per your requirements
-   Default values for env variables are set as follows in the file config/env/env.go  

    ```go
    	EnvVariables = EnvVar{
		RedisAddr: "localhost:6379",
		RedisPass: "",
		GrpcPort:  ":8080",
		MongoUri:  "mongodb://localhost:27017",
	}
    ```
-   Set the env value if it is different from default env values. Following are the env variables getting used in the service -  
    ```
    REDIS_ADDR - Redis server address
    REDIS_PASS - Redis password
    GRPC_PORT  - Port for the gRPC server
    MONGO_URI  - Mongo URI for the MongoDB server
    ```  
-   The service is using redisDB by default, to change it, please read comments in file api/grpcapi/grpc.go

## Roadmap
-   Implememtation of REST and GraphQL sever.
-   Add sample connection with other microservice.

## Owner
Ganeshdip Dumbare  
ganeshdip.dumbare@gmail.com