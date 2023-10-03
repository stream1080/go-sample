package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go-sample/global"
	"go-sample/router"

	"go.uber.org/zap"
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

	global.InitConfig()

	global.InitLogger()

	global.InitMySQL()

	global.InitRedis()

	r := router.Init()

	go func() {
		err := r.Run(fmt.Sprintf(":%d", global.Conf.ServerConfig.Port))
		if err != nil && err != http.ErrServerClosed {
			zap.S().Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// TODO
	zap.S().Info("Shutdown Server ...")
}
