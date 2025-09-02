package controller

import (
	"github.com/codecodesam/mg/micro/auth/api/request"
	"github.com/codecodesam/mg/pkg/base"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Login
// try to get the email and password to login
func Login(ctx *gin.Context) {
	// bind body
	req := &request.LoginReq{}
	// handle error
	err := ctx.MustBindWith(req, binding.JSON)
	if err != nil {
		ctx.JSON(base.ERR_LOGIN_FAIL.HttpCode,
			base.ErrorResponse(base.ERR_LOGIN_FAIL.Code, base.ERR_LOGIN_FAIL.Msg))
	}
	// TODO
}

// Register
// register with email
func Register(ctx *gin.Context) {
	// TODO
}
