package cntechkitgoredis

import (
	"context"
	"fmt"

	gokit "github.com/cntech-io/cntechkit-go"
	"github.com/redis/go-redis/v9"
)

type RedisEnvName string

const (
	REDISDB_HOST     RedisEnvName = "REDISDB_HOST"
	REDISDB_PORT     RedisEnvName = "REDISDB_PORT"
	REDISDB_PASSWORD RedisEnvName = "REDISDB_PASSWORD"
)

type RedisKit struct {
	context context.Context
	client  *redis.Client
}

var env = NewRedisEnv()

func NewRedis() *RedisKit {
	return &RedisKit{
		context: context.Background(),
	}
}

func (rdb *RedisKit) Connect() *RedisKit {
	if env.Host == "" {
		panic("Redis host is empty!")
	}

	if env.Port == "" {
		panic("Redis port is empty!")
	}

	connectionString := fmt.Sprintf("%s:%v", env.Host, env.Port)
	rdb.client = redis.NewClient(&redis.Options{
		Addr: connectionString,
	})

	status := rdb.client.Ping(rdb.context)
	if status.Err() != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", status.Err()))
	}
	gokit.NewLogger(&gokit.LoggerConfig{
		AppName: "cntechkit-goredis",
	}).Info("Connected to Redis!")

	return rdb
}

func (rdb *RedisKit) Do() *redis.Client {
	return rdb.client
}
