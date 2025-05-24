package main

import (
	"micro-server/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	// 创建路由
	r := gin.Default()

	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	routers.AuthRouters(r)

	r.Run(":8111")
}
