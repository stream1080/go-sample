package e

type PError struct {
	Code int
	Msg  string
}

func New(code int, msg string) *PError {
	return &PError{Code: code, Msg: msg}
}

func (p *PError) Error() string {
	return p.Msg
}

var (
	Ok              = New(200, "成功")
	Redirect        = New(301, "请求重定向")
	InvalidArgs     = New(400, "参数错误")
	Unauthorized    = New(401, "身份验证失败")
	Forbidden       = New(403, "未授权的操作")
	NotFound        = New(404, "未找到")
	Conflict        = New(409, "请求冲突")
	TooManyRequests = New(429, "请求太频繁")
	ServerError     = New(500, "系统繁忙")
	DatabaseError   = New(598, "数据库错误")
)
