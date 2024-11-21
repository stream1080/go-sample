package global

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
		zap.S().Error("gorm init error: ", err)
	}
}
