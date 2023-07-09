package router

import (
	"gin-demo/controller"
	_ "gin-demo/docs"
	"gin-demo/middlewares"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 公共方法
	userApi := &controller.UserApi{}
	r.POST("/login", userApi.Login)
	r.POST("/register", userApi.Register)
	r.POST("/send/code", userApi.SendCode)

	authLogin := r.Group("/user", middlewares.AuthLogin())
	authLogin.GET("/info", userApi.GetUserInfo)

	return r
}
