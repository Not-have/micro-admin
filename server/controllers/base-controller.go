package controllers

import "github.com/gin-gonic/gin"

type Base struct{}

func (con Base) success(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  msg,
	})
}

func (con Base) fail(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"code": 500,
		"msg":  msg,
	})
}
