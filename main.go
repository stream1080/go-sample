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

func main() {

	global.Init()

	r := router.Init()

	go func() {
		err := r.Run(fmt.Sprintf(":%d", global.Conf.ServerConfig.Port))
		if err != nil && err != http.ErrServerClosed {
			zap.S().Panicf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.S().Info("Shutdown Server ...")
}
