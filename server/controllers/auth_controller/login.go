package controller

import (
	"fmt"
	"micro-server/models"
	"net/http"

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
		// 查询数据库
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  http.StatusText(http.StatusOK),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  http.StatusText(http.StatusOK),
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
