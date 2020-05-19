package env

import "github.com/ganeshdipdumbare/goenv"

type EnvVar struct {
	RedisAddr string `json:"redis_addr"`
	RedisPass string `json:"redis_pass"`
	GrpcPort  string `json:"grpc_port"`
	MongoUri  string `json:"mongo_uri"`
}

var (
	EnvVariables = EnvVar{
		RedisAddr: "localhost:6379",
		RedisPass: "",
		GrpcPort:  ":8080",
		MongoUri:  "mongodb://localhost:27017",
	}
)

func init() {
	goenv.SyncEnvVar(&EnvVariables)
}
