package middlewares

import (
	"gin-demo/controller"
	"gin-demo/ulits"

	"github.com/gin-gonic/gin"
)

// AuthLogin 登录校验中间件
func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		user, err := ulits.AnalyseToken(auth)
		if err != nil || user == nil {
			c.Abort()
			controller.ResponseError(c, controller.Unauthorized)
			return
		}
		c.Next()
	}
}
