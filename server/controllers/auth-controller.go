package controllers

import (
	"encoding/json"
	"fmt"
	"micro-server/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Base
}

func (con LoginController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	captchaId := c.PostForm("captchaId")
	verifyValue := c.PostForm("verifyValue")

	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag {

		// 1. 根据用户名查用户（用户名是唯一索引）
		userinfoList := []models.AdminUser{}

		err := models.DB.Where("username=? AND password=?", username, password).Find(&userinfoList).Error

		if err != nil {
			// 模糊提示，避免暴露用户是否存在的信息
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "用户名或密码错误",
			})
			return
		}

		if len(userinfoList) > 0 {
			//3、执行登录 保存用户信息 执行跳转
			session := sessions.Default(c)
			//注意：session.Set没法直接保存结构体对应的切片 把结构体转换成json字符串
			userinfoSlice, _ := json.Marshal(userinfoList[0])
			session.Set("userinfo", string(userinfoSlice))
			session.Save()
		}

		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  http.StatusText(http.StatusOK),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "验证码错误",
		})
	}
}

func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := models.MakeCaptcha()

	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}
