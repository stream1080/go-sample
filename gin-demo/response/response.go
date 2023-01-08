package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}
func Error(code int, msg string) Response {
	return Response{
		code,
		msg,
		nil,
	}
}

func Ok(c *gin.Context) {
	OkWithData(c, nil)
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(c, SUCCESS.Code, SUCCESS.Msg, data)
}

func OkWithMsg(c *gin.Context, msg string) {
	Result(c, SUCCESS.Code, msg, nil)
}

func Fail(c *gin.Context, err Response) {
	Result(c, err.Code, err.Msg, nil)
}

func FailWithMsg(c *gin.Context, err Response, msg string) {
	Result(c, err.Code, msg, nil)
}
