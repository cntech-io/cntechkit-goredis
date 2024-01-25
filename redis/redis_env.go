package redis

import (
	e "github.com/cntech-io/cntechkit-go/v2/env"
	"github.com/cntech-io/cntechkit-go/v2/logger"
	"github.com/joho/godotenv"
)

type RedisEnvName string

const (
	REDISDB_HOST     RedisEnvName = "REDISDB_HOST"
	REDISDB_PORT     RedisEnvName = "REDISDB_PORT"
	REDISDB_PASSWORD RedisEnvName = "REDISDB_PASSWORD"
)

type RedisEnv struct {
	Host     string
	Port     string
	Password string
}

func NewRedisEnv() *RedisEnv {
	if err := godotenv.Load(); err != nil {
		logger.NewLogger(&logger.LoggerConfig{
			AppName: "cntechkit-goredis",
		}).Info("Failed to load .env file!")
	}
	return &RedisEnv{
		Host:     e.GetString(string(REDISDB_HOST), false),
		Port:     e.GetString(string(REDISDB_PORT), false),
		Password: e.GetString(string(REDISDB_PASSWORD), false),
	}
}
