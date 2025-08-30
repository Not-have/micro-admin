package routers

import (
	"micro-server/controllers"
	"micro-server/middlewares"

	"github.com/gin-gonic/gin"
)

func RoleInit(r *gin.Engine) {
	user := r.Group("/role", middlewares.InitAdminAuthMiddleware)

	// 用户相关路由
	{
		user.POST("/create", controllers.RoleController{}.RoleCreate)
		user.GET("/list", controllers.RoleController{}.RoleList)
	}
}
