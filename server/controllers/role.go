package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	Base
}

func (con RoleController) RoleList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  http.StatusText(http.StatusOK),
	})
}

func (con RoleController) RoleCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  http.StatusText(http.StatusOK),
	})
}
