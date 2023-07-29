package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, code Code, msg interface{}, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg:": msg,
		"data": data,
	})
}

func ResponseOK(c *gin.Context, data interface{}) {
	Response(c, SUCCESS, SUCCESS.Msg(), data)
}

func ResponseError(c *gin.Context, code Code) {
	Response(c, code, code.Msg(), nil)
}

func ResponseErrorWith(c *gin.Context, code Code, msg interface{}) {
	Response(c, code, msg, nil)
}
