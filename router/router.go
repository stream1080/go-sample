package router

import (
	"fmt"

	"go-sample/controller"
	_ "go-sample/docs"
	"go-sample/middlewares"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.NoRoute(func(c *gin.Context) {
		msg := fmt.Sprintf("not found: %s %s", c.Request.Method, c.Request.RequestURI)
		controller.ResponseErrorWithMsg(c, controller.NotFound, msg)
	})

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
