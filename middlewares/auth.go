package middlewares

import (
	"strings"

	"go-sample/consts"
	"go-sample/pkg/jwt"
	"go-sample/pkg/response"

	"github.com/gin-gonic/gin"
)

// Auth 登录校验中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := c.GetHeader(consts.AuthHeader)
		parts := strings.SplitN(auth, " ", 2)
		if !(len(parts) == 2 && parts[0] == consts.Bearer) {
			response.Error(c, response.Unauthorized)
			c.Abort()
			return
		}

		user, err := jwt.AnalyseToken(parts[1], "")
		if err != nil {
			response.Error(c, response.Unauthorized)
			c.Abort()
			return
		}

		c.Set(consts.CtxUserKey, user)
		c.Next()
	}
}
