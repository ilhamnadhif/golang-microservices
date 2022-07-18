package app

import (
	"github.com/go-redis/redis/v8"
	"user/config"
)

func InitRedis(config config.RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.HostPort,
		Password: config.Password,
		DB:       config.DbNumber,
	})
	return rdb
}
