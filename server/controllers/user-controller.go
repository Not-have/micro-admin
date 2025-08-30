package controllers

import (
	"encoding/json"
	"fmt"
	"micro-server/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserInfoController struct {
	Base
}

func (con UserInfoController) UserInfo(c *gin.Context) {

	session := sessions.Default(c)
	//注意：session.Set没法直接保存结构体对应的切片 把结构体转换成json字符串
	userinfoStr, ok := session.Get("userinfo").(string)
	fmt.Println(userinfoStr)

	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusNonAuthoritativeInfo,
			"msg":  "获取用户信息失败",
		})
		return
	}

	// 注意：这里应该解析为单个AdminUser对象，而不是切片
	var user models.AdminUser

	// 字符串转为结构体
	err := json.Unmarshal([]byte(userinfoStr), &user)

	if err != nil {
		// 输出详细错误和原始数据，方便调试
		c.JSON(200, gin.H{"error": "解析用户信息失败: " + err.Error()})
		return
	}

	// 成功解析后，即可使用用户信息
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  http.StatusText(http.StatusOK),
		"data": user,
	})
}
