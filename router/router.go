package router

import (
	_ "demo/docs"
	"demo/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 路由规则
	r.GET("/", service.Index)

	r.GET("/user-info", service.GetUserInfo)

	r.POST("/user-login", service.Login)

	r.POST("/user/register", service.Register)

	return r
}
