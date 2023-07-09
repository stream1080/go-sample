package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code RespCode    `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(c *gin.Context, code RespCode, msg string, data ...interface{}) {
	c.JSON(http.StatusOK, Result{
		code,
		msg,
		data,
	})
}

func ResponseOK(c *gin.Context, data ...interface{}) {
	Response(c, OK, OK.Msg(), data)
}

func ResponseError(c *gin.Context, code RespCode) {
	Response(c, code, code.Msg(), nil)
}
