package service

import (
	"gin-demo/api"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	res := api.NewResult(c)
	res.Success(nil)
}
