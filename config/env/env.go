package env

import "github.com/ganeshdipdumbare/goenv"

type EnvVar struct {
	RedisAddr string `json:"redis_addr"`
	RedisPass string `json:"redis_pass"`
	GrpcPort  string `json:"grpc_port"`
}

var (
	EnvVariables = EnvVar{
		RedisAddr: "localhost:6379",
		RedisPass: "",
		GrpcPort:  ":8080",
	}
)

func init() {
	goenv.SyncEnvVar(&EnvVariables)
}
