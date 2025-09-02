package base

type CustomizeError struct {
	HttpCode int
	Code     int
	Msg      string
}

var (
	ERR_LOGIN_FAIL = CustomizeError{HttpCode: 400, Code: 0, Msg: "login fail(reason: email or password is wrong)"}
)
