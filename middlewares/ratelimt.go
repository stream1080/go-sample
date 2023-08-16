package middlewares

import (
	"time"

	"go-sample/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimit(interval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(interval, cap)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) == 0 {
			response.Error(c, response.TooManyRequests)
			c.Abort()
			return
		}

		c.Next()
	}
}
