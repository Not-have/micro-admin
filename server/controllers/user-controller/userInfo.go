package controllers

import (
	"encoding/json"
	BaseController "micro-server/controllers/base-controller"
	"micro-server/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserInfoController struct {
	BaseController.BaseController
}

func (con UserInfoController) UserInfo(c *gin.Context) {

	session := sessions.Default(c)
	//注意：session.Set没法直接保存结构体对应的切片 把结构体转换成json字符串
	userinfoStr, ok := session.Get("userinfo").(string)

	if !ok {
		c.JSON(200, gin.H{"error": "未找到用户信息"})
		return
	}

	var userList []models.AdminUser
	err := json.Unmarshal([]byte(userinfoStr), &userList)
	if err != nil {
		c.JSON(200, gin.H{"error": "解析用户信息失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  http.StatusText(http.StatusOK),
		"data": userList,
	})
}
