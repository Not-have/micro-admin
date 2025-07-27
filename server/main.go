package main

import (
	"micro-server/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	// 创建路由
	r := gin.Default()

	// 关键步骤：初始化session中间件
	// 使用cookie存储session数据，"secret"是加密密钥（生产环境需更换为复杂密钥）
	store := cookie.NewStore([]byte("secret"))

	//配置session的中间件 store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("micro-admin", store))

	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	routers.AuthRoutersInit(r)
	routers.UserInit(r)

	r.Run(":8111")
}
