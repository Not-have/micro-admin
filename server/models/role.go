// 操作数据库

package models

import "time"

// 角色表
type Role struct {
	// 角色ID， 主键，自增
	RoleId uint `gorm:"column:role_id;primaryKey;autoIncrement" json:"role_id"`
	// 角色名称
	RoleName   string    `gorm:"column:role_name" json:"role_name"`
	RoleDesc   string    `gorm:"column:role_desc" json:"role_desc"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
}

// 返回给前端的角色列表
type RoleList struct {
	RoleId   uint   `json:"role_id"`   // 角色ID
	RoleName string `json:"role_name"` // 角色名称
	RoleDesc string `json:"role_desc"` // 角色描述
}

func (Role) TableName() string {
	// 指定表名
	return "role"
}
