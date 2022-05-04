package middlewares

import (
	"gin-demo/api"
	"gin-demo/ulits"

	"github.com/gin-gonic/gin"
)

// AuthLogin 登录校验中间件
func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := api.NewResult(c)
		auth := c.GetHeader("Authorization")
		user, err := ulits.AnalyseToken(auth)
		if err != nil || user == nil {
			c.Abort()
			res.Error(api.Forbidden)
			return
		}
		c.Next()
	}
}
