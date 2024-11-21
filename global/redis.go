package global

import (
	"fmt"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func InitRedis() {

	RDB = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", Conf.RedisConfig.Host, Conf.RedisConfig.Port),
		Password:     Conf.RedisConfig.Password,
		DB:           Conf.RedisConfig.DB,
		PoolSize:     Conf.RedisConfig.PoolSize,
		MinIdleConns: Conf.RedisConfig.MinIdleConns,
	})

	_, err := RDB.Ping().Result()
	if err != nil {
		zap.S().Error("redis init error: ", err)
	}
}
