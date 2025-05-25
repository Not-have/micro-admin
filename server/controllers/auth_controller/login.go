package controller

import (
	"fmt"
	"micro-server/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	BaseController "micro-server/controllers/base_controller"
)

type LoginController struct {
	BaseController.BaseController
}

func (con LoginController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	captchaId := c.PostForm("captchaId")
	verifyValue := c.PostForm("verifyValue")

	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag {

		fmt.Println(username)

		// 1. 根据用户名查用户（用户名是唯一索引）
		var user models.AdminUser

		err := models.DB.Where("username=? AND password=?", username, password).Find(&user).Error

		if err != nil {
			// 模糊提示，避免暴露用户是否存在的信息
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "用户名或密码错误",
			})
			return
		}

		//2、执行登录 保存用户信息 执行跳转
		session := sessions.Default(c)
		session.Set("user", user)
		session.Save()

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  http.StatusText(http.StatusOK),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
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
