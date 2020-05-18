package main

import (
	"hex-gopher/api/grpcapi"
	"hex-gopher/config/env"
)

func main() {
	envVar := &env.EnvVariables
	grpcApiServer := grpcapi.NewServer(envVar)
	grpcApiServer.StartServer()
}
