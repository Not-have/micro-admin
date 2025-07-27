package routers

import (
	controllers "micro-server/controllers/user-controller"
	"micro-server/middlewares"

	"github.com/gin-gonic/gin"
)

func UserInit(r *gin.Engine) {
	user := r.Group("/user", middlewares.InitAdminAuthMiddleware)

	{
		user.GET("/info", controllers.UserInfoController{}.UserInfo)
	}
}
