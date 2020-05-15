package main

import "hex-gopher/api/grpcapi"

func main() {
	grpcApiServer := grpcapi.NewServer()
	grpcApiServer.StartServer()
}
