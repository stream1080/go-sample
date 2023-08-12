package middlewares

import (
	"strings"

	"go-sample/pkg/jwt"
	"go-sample/pkg/response"

	"github.com/gin-gonic/gin"
)

const (
	AuthHeader = "Authorization"
	Bearer     = "Bearer"
	CtxUserKey = "user"
)

// Auth 登录校验中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := c.GetHeader(AuthHeader)
		parts := strings.SplitN(auth, " ", 2)
		if !(len(parts) == 2 && parts[0] == Bearer) {
			response.Error(c, response.Unauthorized)
			c.Abort()
			return
		}

		user, err := jwt.AnalyseToken(parts[1])
		if err != nil {
			response.Error(c, response.Unauthorized)
			c.Abort()
			return
		}

		c.Set(CtxUserKey, user)
		c.Next()
	}
}
