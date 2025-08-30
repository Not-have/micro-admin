package routers

import (
	"micro-server/controllers"

	"github.com/gin-gonic/gin"
)

func RoleInit(r *gin.Engine) {
	user := r.Group("/role")

	// 用户相关路由
	{
		user.POST("/create", controllers.RoleController{}.RoleCreate)
		user.GET("/list", controllers.RoleController{}.RoleList)
	}
}
