package base

type CommonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Tid  string      `json:"tid"`
	Data interface{} `json:"data"`
}

func Wrap(code int, msg string, data interface{}) *CommonResponse {
	return &CommonResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func ErrorResponse(code int, msg string) *CommonResponse {
	return &CommonResponse{
		Code: code,
		Msg:  msg,
	}
}
