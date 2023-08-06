package middlewares

import (
	"go-sample/controller"
	"go-sample/pkg/ulits"

	"github.com/gin-gonic/gin"
)

// Auth 登录校验中间件
func Auth() gin.HandlerFunc {
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
