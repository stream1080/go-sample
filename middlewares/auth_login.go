package middlewares

import (
	"demo/api"
	"demo/ulits"

	"github.com/gin-gonic/gin"
)

// AuthAdminCheck is a middleware function that checks if the user is authenticated with admin role.
func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := api.NewResult(c)
		auth := c.GetHeader("Authorization")
		user, err := ulits.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			res.Error(api.Forbidden)
			return
		}
		if user == nil {
			c.Abort()
			res.Error(api.Forbidden)
			return
		}
		c.Next()
	}
}
