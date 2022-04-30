package router

import (
	_ "demo/docs"
	"demo/middlewares"
	"demo/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 公共方法
	r.GET("/", service.Index)
	r.POST("/login", service.Login)
	r.POST("/register", service.Register)
	r.POST("/send/code", service.SendCode)

	authLogin := r.Group("/user", middlewares.AuthLogin())
	authLogin.GET("/info", service.GetUserInfo)

	return r
}
