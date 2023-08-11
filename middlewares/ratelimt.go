package middlewares

import (
	"go-sample/controller"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimit(interval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(interval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就返回响应
		if bucket.TakeAvailable(1) == 0 {
			controller.ResponseError(c, controller.TooManyRequests)
			c.Abort()
			return
		}
		// 取到令牌就放行
		c.Next()
	}
}
