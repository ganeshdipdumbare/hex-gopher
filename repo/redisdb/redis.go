package redisdb

import (
	"errors"
	"hex-gopher/app"

	"github.com/go-redis/redis"
	errs "github.com/pkg/errors"
)

type RedisDB struct {
	client *redis.Client
}

func getRedisClient(addr string, pwd string) (*redis.Client, error) {
	cl := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB:       0,   // use default DB
	})
	_, err := cl.Ping().Result()
	if err != nil {
		return nil, errs.Wrap(err, "failed to create client for redis")
	}
	return cl, nil
}

func NewRedisDB(addr string, pwd string) (*RedisDB, error) {
	cl, err := getRedisClient(addr, pwd)
	if err != nil {
		return nil, errs.Wrap(err, "failed to get client for redis")
	}

	return &RedisDB{
		client: cl,
	}, nil
}
func (r *RedisDB) SaveGopher(g *app.Gopher) (string, error) {
	if g == nil {
		return "", errs.Wrap(errors.New("nil gopher passed"), "nil gopher passed")
	}

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
