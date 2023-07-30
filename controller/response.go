package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, data interface{}) {
	Response(c, SUCCESS, SUCCESS.Msg(), data)
}

func ResponseError(c *gin.Context, code Code) {
	Response(c, code, code.Msg(), nil)
}

func ResponseErrorWithMsg(c *gin.Context, code Code, msg interface{}) {
	Response(c, code, msg, nil)
}

func Response(c *gin.Context, code Code, msg interface{}, data interface{}) {

	httpCode := http.StatusOK
	if code >= 100 && code < 599 {
		httpCode = int(code)
	}

	c.JSON(httpCode, gin.H{
		"code": code,
		"msg:": msg,
		"data": data,
	})
}
