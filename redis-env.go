package cntechkitgoredis

import (
	gokit "github.com/cntech-io/cntechkit-go"
	"github.com/cntech-io/cntechkit-go/utils"
	"github.com/joho/godotenv"
)

type RedisEnv struct {
	Host     string
	Port     string
	Password string
}

func NewRedisEnv() *RedisEnv {
	if err := godotenv.Load(); err != nil {
		gokit.NewLogger(&gokit.LoggerConfig{
			AppName: "cntechkit-goredis",
		}).Info("Failed to load .env file!")
	}
	return &RedisEnv{
		Host:     utils.GetStringEnv(string(REDISDB_HOST), false),
		Port:     utils.GetStringEnv(string(REDISDB_PORT), false),
		Password: utils.GetStringEnv(string(REDISDB_PASSWORD), false),
	}
}
