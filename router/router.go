package router

import (
	"embed"
	"fmt"
	"mime"
	"net/http"
	"strings"

	"go-sample/controller"
	_ "go-sample/docs"
	"go-sample/middlewares"
	"go-sample/pkg/response"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Handler(fs embed.FS) *gin.Engine {

	r := gin.New()

	r.Use(middlewares.Cors(), middlewares.Logger(), middlewares.Recovery(true))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/api/") {
			msg := fmt.Sprintf("not found: %s %s", c.Request.Method, c.Request.RequestURI)
			response.WithMsg(c, response.NotFound, msg)
			return
		}

		// 读取文件内容
		if data, err := fs.ReadFile(fmt.Sprintf("dist%s", path)); err != nil {
			// 如果文件不存在，返回首页 index.html
			if data, err = fs.ReadFile("dist/index.html"); err != nil {
				response.WithMsg(c, response.NotFound, err.Error())
			} else {
				c.Data(http.StatusOK, mime.TypeByExtension(".html"), data)
			}
		} else {
			// 如果文件存在，根据请求的文件后缀，设置正确的mime type，并返回文件内容
			s := strings.Split(path, ".") // 分割路径，获取文件后缀
			c.Data(http.StatusOK, mime.TypeByExtension(fmt.Sprintf(".%s", s[len(s)-1])), data)
		}
	})

	api := r.Group("/api")
	v1 := api.Group("/v1")

	// 公共方法
	userApi := &controller.UserApi{}
	v1.POST("/login", userApi.Login)
	v1.POST("/register", userApi.Register)
	v1.POST("/send/code", userApi.SendCode)

	authLogin := v1.Group("/user", middlewares.Auth())
	authLogin.GET("/info", userApi.UserInfo)

	return r
}
