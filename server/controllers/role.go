package controllers

import (
	"fmt"
	"micro-server/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	Base
}

func (con RoleController) RoleList(c *gin.Context) {

	roleList := []models.Role{}

	err := models.DB.Find(&roleList).Error

	var responseList []models.RoleList
	for _, role := range roleList {
		responseList = append(responseList, models.RoleList{
			RoleId:   role.RoleId,
			RoleName: role.RoleName,
			RoleDesc: role.RoleDesc,
		})
	}

	if err != nil {
		con.Base.fail(c, "获取角色列表失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  http.StatusText(http.StatusOK),
		"data": responseList,
	})
}

func (con RoleController) RoleCreate(c *gin.Context) {
	role_name := strings.Trim(c.PostForm("role_name"), " ")
	role_desc := strings.Trim(c.PostForm("role_desc"), " ")

	fmt.Println(role_name, role_desc)

	if role_name == "" || role_desc == "" {
		con.Base.fail(c, "参数错误")
	}

	role := models.Role{}

	role.RoleName = role_name
	role.RoleDesc = role_desc
	role.CreateTime = time.Now()

	// 来个检查，禁止重名
	var count int64
	models.DB.Model(&models.Role{}).Where("role_name = ?", role_name).Count(&count)
	if count > 0 {
		con.Base.fail(c, "角色名已存在")
		return
	}

	// 链接数据库
	err := models.DB.Create(&role).Error
	if err != nil {
		con.Base.fail(c, "创建失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  http.StatusText(http.StatusOK),
	})
}

func (con RoleController) RoleDelete(c *gin.Context) {
	id := c.PostForm("id")

	if id == "" {
		con.Base.fail(c, "参数错误")
		return
	}

	result := models.DB.Delete(&models.Role{}, id)

	if result.RowsAffected == 0 {
		con.Base.fail(c, "角色不存在")
		return
	}

	if result.Error != nil {
		con.Base.fail(c, "删除失败")
		return
	}

	con.Base.success(c, "删除成功")
}
