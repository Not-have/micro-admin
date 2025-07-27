package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitAdminAuthMiddleware(c *gin.Context) {

	//进行权限判断 没有登录的用户 不能进入后台管理中心
	fmt.Println("权限判断")
}
