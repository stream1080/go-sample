package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-redis/redis"
)

var (
	DB    = InitMySQL()
	Redis = InitRedis()
)

func InitMySQL() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error : ", err)
	}
	return db
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
