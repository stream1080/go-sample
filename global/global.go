package global

import (
	"fmt"
	"log"
	"os"

	"go-sample/config"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	data, err := os.ReadFile("./app.yaml")
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

func InitLogger() {
	zc := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     getEncoder(),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}
	lg, err := zc.Build()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(lg)

	zap.S().Info("init logger successful!")
}

func getEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "name",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
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
