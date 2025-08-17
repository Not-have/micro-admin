package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/**
 * 最为底层的权限判断
 */
func InitAdminAuthMiddleware(c *gin.Context) {

	//进行权限判断 没有登录的用户 不能进入后台管理中心
	fmt.Println("权限判断")

	pathName := strings.Split(c.Request.URL.Path, "?")

	fmt.Println(pathName)
	// 获取 Session 里面的信息

	sessions := sessions.Default(c)

	userInfo := sessions.Get("userinfo")

	_, ok := userInfo.(string)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "请登录",
		})

		// 中间件提前结束
		c.Abort()
	}
}
