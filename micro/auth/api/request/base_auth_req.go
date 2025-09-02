package request

type LoginReq struct {
	Email string `json:"email" binding:"required"`
	Pwd   string `json:"pwd" binding:"required"`
}
