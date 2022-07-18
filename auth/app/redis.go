package app

import (
	"auth/config"
	"github.com/go-redis/redis/v8"
)

func InitRedis(config config.RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.HostPort,
		Password: config.Password,
		DB:       config.DbNumber,
	})
	return rdb
}
