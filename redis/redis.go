package redis

import (
	"context"
	"fmt"

	"github.com/cntech-io/cntechkit-go/v2/logger"
	"github.com/redis/go-redis/v9"
)

type redisInstance struct {
	context context.Context
	client  *redis.Client
}

var env = NewRedisEnv()

func NewRedis() *redisInstance {
	return &redisInstance{
		context: context.Background(),
	}
}

func (rdb *redisInstance) Connect() *redisInstance {
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
	logger.NewLogger(&logger.LoggerConfig{
		AppName: "cntechkit-goredis",
	}).Info("Connected to Redis!")

	return rdb
}

func (rdb *redisInstance) Do() *redis.Client {
	return rdb.client
}
