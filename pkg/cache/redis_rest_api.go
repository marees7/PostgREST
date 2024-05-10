package cache

import (
	"context"
	"postgrest/config"

	"github.com/go-redis/redis/v8"
)

type RedisRestAPI struct {
	Client0 *redis.Client
}

var RedisRestAPIClients *RedisRestAPI

func NewRedisRestAPIConnection() error {
	redisOpt0 := &redis.Options{
		Network:  config.Configuration.Get("REDIS_PROTOCOL"),
		Addr:     config.Configuration.Get("REDIS_HOST"),
		Password: config.Configuration.Get("REDIS_PASSWORD"),
		DB:       0,
	}
	client0 := redis.NewClient(redisOpt0)
	if _, err := client0.Ping(context.Background()).Result(); err != nil {
		return err
	}
	RedisRestAPIClients = &RedisRestAPI{
		Client0: client0,
	}
	return nil
}
