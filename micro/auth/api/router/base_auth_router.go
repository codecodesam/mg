package router

import (
	"github.com/codecodesam/mg/micro/auth/api/controller"
	"github.com/gin-gonic/gin"
)

// Register
// register router
func Register(engine *gin.Engine) {
	// define the group
	group := engine.Group("/base_auth")
	// create some http method mapping
	group.POST("/login", controller.Login)
	group.POST("/register", controller.Register)
}
