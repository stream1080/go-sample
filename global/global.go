package global

import (
	"fmt"
	"log"
	"os"

	"go-sample/config"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Conf *config.Config
	DB   *gorm.DB
	RDB  *redis.Client
)

func InitConfig() {
	data, err := os.ReadFile("./config/app.yaml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &Conf); err != nil {
		panic(err)
	}

	if Conf.ServerConfig.Mode == "" {
		Conf.ServerConfig.Mode = gin.DebugMode
	}
}

func InitMySQL() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Conf.MySQLConfig.User,
		Conf.MySQLConfig.Password,
		Conf.MySQLConfig.Host,
		Conf.MySQLConfig.Port,
		Conf.MySQLConfig.DB)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error : ", err)
	}
}

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
		log.Println("redis Init Error : ", err)
	}
}
