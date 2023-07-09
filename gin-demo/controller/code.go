package controller

type RespCode int

const (
	OK              RespCode = 200
	NeedRedirect    RespCode = 301
	InvalidArgs     RespCode = 400
	Unauthorized    RespCode = 401
	Forbidden       RespCode = 403
	NotFound        RespCode = 404
	Conflict        RespCode = 409
	TooManyRequests RespCode = 429
	ServerError     RespCode = 500

	CodeExpire RespCode = 5001
	CodeError  RespCode = 5002
	UserExist  RespCode = 5003
)

var codeMsg = map[RespCode]string{
	OK:              "ok",
	NeedRedirect:    "need redirect",
	InvalidArgs:     "invalid params",
	Unauthorized:    "unauthorized",
	Forbidden:       "forbidden",
	NotFound:        "not found",
	Conflict:        "conflict",
	TooManyRequests: "too many requests",
	ServerError:     "server error",

	CodeExpire: "code expire",
	CodeError:  "code error",
	UserExist:  "user exist",
}

func (r RespCode) Msg() string {
	msg, ok := codeMsg[r]
	if ok {
		return msg
	}

	return codeMsg[ServerError]
}
