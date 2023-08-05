package global

import (
	"os"

	"go-sample/config"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

var (
	Conf *config.Config
)

func InitConfig() {
	data, err := os.ReadFile("./config/app.yaml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, Conf); err != nil {
		panic(err)
	}

	if Conf.ServerConfig.Mode == "" {
		Conf.ServerConfig.Mode = gin.DebugMode
	}
}
