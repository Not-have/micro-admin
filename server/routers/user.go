package routers

import (
	"micro-server/controllers"
	"micro-server/middlewares"

	"github.com/gin-gonic/gin"
)

func UserInit(r *gin.Engine) {
	user := r.Group("/user", middlewares.InitAdminAuthMiddleware)

	// 用户相关路由
	{
		user.GET("/info", controllers.UserInfoController{}.UserInfo)
	}
}
