package routers

import (
	"micro-server/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutersInit(r *gin.Engine) {
	r.POST("/login", controllers.LoginController{}.Login)
	r.GET("/captcha", controllers.LoginController{}.Captcha)
}
