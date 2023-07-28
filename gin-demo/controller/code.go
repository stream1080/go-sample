package controller

type Code int

const (
	SUCCESS         Code = 200
	Redirect        Code = 301
	InvalidArgs     Code = 400
	Unauthorized    Code = 401
	Forbidden       Code = 403
	NotFound        Code = 404
	Conflict        Code = 409
	TooManyRequests Code = 429
	ServerError     Code = 500

	_ = 1000 + iota
	CodeExpire
	CodeError
	UserExist
)

var codeMsg = map[Code]string{
	SUCCESS:         "success",
	Redirect:        "redirect",
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

func (c Code) Msg() string {
	msg, ok := codeMsg[c]
	if ok {
		return msg
	}

	return codeMsg[ServerError]
}
