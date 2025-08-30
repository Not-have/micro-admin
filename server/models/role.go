// 操作数据库

package models

import "time"

// 角色表
type Role struct {
	RoleId     uint      `gorm:"column:role_id;primaryKey;autoIncrement" json:"role_id"`
	RoleName   string    `gorm:"column:role_name" json:"role_name"`
	RoleDesc   string    `gorm:"column:role_desc" json:"role_desc"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
}

func (Role) TableName() string {
	// 指定表名
	return "role"
}
