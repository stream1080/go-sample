package models

import (
	"github.com/go-redis/redis"
)

var (
	Redis = InitRedis()
)

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
