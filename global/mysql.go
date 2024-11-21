package global

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMySQL() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Conf.MySQLConfig.User,
		Conf.MySQLConfig.Password,
		Conf.MySQLConfig.Host,
		Conf.MySQLConfig.Port,
		Conf.MySQLConfig.DB)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		zap.L().Panic("gorm init error: ", zap.Error(err))
	}

	sqlDB, err := DB.DB()
	if err != nil {
		zap.L().Panic("get sqlDB error", zap.Error(err))
	}

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)
}
