package global

import (
	"log"

	"go-sample/config"

	"github.com/Netflix/go-env"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	Conf config.Config
	DB   *gorm.DB
	RDB  *redis.Client
)

func InitConfig() {
	if _, err := env.UnmarshalFromEnviron(&Conf); err != nil {
		log.Panicf("init config failed with %s\n", err)
	}
}

func Init() {
	InitConfig()
	InitLogger()
	InitMySQL()
	InitRedis()
}
