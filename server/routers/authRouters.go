package routers

import (
	controller "micro-server/controllers/auth-controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutersInit(r *gin.Engine) {
	r.POST("/login", controller.LoginController{}.Login)
	r.GET("/captcha", controller.LoginController{}.Captcha)

}
