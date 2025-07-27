package BaseController

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (con BaseController) success(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  msg,
	})
}

func (con BaseController) fail(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"code": 500,
		"msg":  msg,
	})
}
