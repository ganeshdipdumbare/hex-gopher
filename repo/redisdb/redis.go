package redisdb

import (
	"hex-gopher/app"

	"github.com/go-redis/redis"
	errs "github.com/pkg/errors"
)

type RedisDB struct {
	client *redis.Client
}

func NewRedisDB(addr string, pwd string) (*RedisDB, error) {
	cl := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB:       0,   // use default DB
	})
	_, err := cl.Ping().Result()
	if err != nil {
		return nil, errs.Wrap(err, "failed to create client for Redis")
	}

	return &RedisDB{
		client: cl,
	}, nil
}
func (r *RedisDB) SaveGopher(g *app.Gopher) (string, error) {
	err := r.client.Set(g.Id, g.Name, 0).Err()
	if err != nil {
		return "", errs.Wrap(err, "unable to save the gopher to redis")
	}
	return g.Id, nil
}

func (r *RedisDB) GetGopher(id string) (*app.Gopher, error) {
	gopherName, err := r.client.Get(id).Result()
	if err != nil {
		return nil, errs.Wrap(err, "unable to save the gopher to redis")
	}
	return &app.Gopher{
		Id:   id,
		Name: gopherName,
	}, nil
}
