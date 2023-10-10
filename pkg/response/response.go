package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code Code        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	Response(c, SUCCESS, SUCCESS.Msg(), data)
}

func Error(c *gin.Context, code Code) {
	Response(c, code, code.Msg(), nil)
}

func WithMsg(c *gin.Context, code Code, msg string) {
	Response(c, code, msg, nil)
}

func Response(c *gin.Context, code Code, msg string, data interface{}) {

	httpCode := http.StatusOK
	if code >= 100 && code < 599 {
		httpCode = int(code)
	}

	c.JSON(httpCode, &Result{code, msg, data})
}
