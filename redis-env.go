package cntechkitgoredis

import (
	"github.com/cntech-io/cntechkit-go/utils"
)

type RedisEnv struct {
	Host     string
	Port     string
	Password string
}

func NewRedisEnv() *RedisEnv {
	return &RedisEnv{
		Host:     utils.GetStringEnv(string(REDISDB_HOST), false),
		Port:     utils.GetStringEnv(string(REDISDB_PORT), false),
		Password: utils.GetStringEnv(string(REDISDB_PASSWORD), false),
	}
}
