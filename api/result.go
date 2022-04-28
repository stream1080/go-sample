package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResultCtx struct {
	Ctx *gin.Context
}

//返回的结果：
type Result struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

func NewResult(ctx *gin.Context) *ResultCtx {
	return &ResultCtx{Ctx: ctx}
}

func NewError(code int, msg string) Result {
	return Result{
		Code: code,
		Msg:  msg,
		Data: gin.H{},
	}
}

func (r *ResultCtx) Success(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, Result{
		Code: 200,
		Data: data,
		Msg:  "",
	})
}

func (r *ResultCtx) Error(result Result) {
	r.Ctx.JSON(http.StatusOK, Result{
		Code: result.Code,
		Data: gin.H{},
		Msg:  result.Msg,
	})
}
