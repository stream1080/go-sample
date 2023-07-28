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

	api := r.Group("/api")
	v1 := api.Group("/v1")

	// 公共方法
	userApi := &controller.UserApi{}
	v1.POST("/login", userApi.Login)
	v1.POST("/register", userApi.Register)
	v1.POST("/send/code", userApi.SendCode)

	authLogin := v1.Group("/user", middlewares.AuthLogin())
	authLogin.GET("/info", userApi.GetUserInfo)

	return r
}
