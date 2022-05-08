package api

var (
	NeedRedirect    = NewError(301, "need redirect")
	InvalidArgs     = NewError(400, "invalid args")
	Unauthorized    = NewError(401, "unauthorized")
	Forbidden       = NewError(403, "forbidden")
	NotFound        = NewError(404, "not found")
	Conflict        = NewError(409, "entry exist")
	TooManyRequests = NewError(429, "too many requests")
	ResultError     = NewError(500, "response result error")
	DatabaseError   = NewError(598, "database err")
	CSRFDetected    = NewError(599, "csrf attack detected")

	UserError  = NewError(5001, "username or password error")
	CodeExpire = NewError(5002, "verification expire")
	CodeError  = NewError(5003, "verification error")
	UserExist  = NewError(5004, "user Exist")
)
