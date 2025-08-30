package routers

import (
	"micro-server/controllers"

	"github.com/gin-gonic/gin"
)

func RoleInit(r *gin.Engine) {
	role := r.Group("/role")

	// 用户相关路由
	{
		role.POST("/create", controllers.RoleController{}.RoleCreate)
		role.GET("/list", controllers.RoleController{}.RoleList)
		role.DELETE("/delete", controllers.RoleController{}.RoleDelete)
	}
}
