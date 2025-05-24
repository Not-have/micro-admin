package routers

import (
	controller "micro-server/controllers/auth_controller"

	"github.com/gin-gonic/gin"
)

func AuthRouters(r *gin.Engine) {
	r.POST("/login", controller.LoginController{}.Login)
}

func CaptchaRouters(r *gin.Engine) {
	r.GET("/captcha", controller.LoginController{}.Captcha)
}
