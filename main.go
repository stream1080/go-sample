package main

import (
	"fmt"

	"go-sample/global"
	"go-sample/router"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://github.com/stream108

// @contact.name 一江溪水
// @contact.url https://github.com/stream1080
// @contact.email example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {

	// 初始化配置
	global.InitConfig()

	global.InitMySQL()

	global.InitRedis()

	r := router.Init()

	r.Run(fmt.Sprintf(":%d", global.Conf.ServerConfig.Port))
}
