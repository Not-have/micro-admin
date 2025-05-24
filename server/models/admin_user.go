package models

import "time"

type AdminUser struct {
	UserId       uint       `gorm:"column:user_id;primaryKey;autoIncrement"`
	Username     string     `gorm:"column:username;uniqueIndex"`
	PasswordHash string     `gorm:"column:password_hash"`
	RealName     string     `gorm:"column:real_name"`
	Gender       string     `gorm:"column:gender"`
	Mobile       string     `gorm:"column:mobile;uniqueIndex"`
	Email        string     `gorm:"column:email;uniqueIndex"`
	DepartmentId int        `gorm:"column:department_id;index"`
	RoleType     string     `gorm:"column:role_type"`
	UserStatus   string     `gorm:"column:user_status;index"`
	CreateTime   time.Time  `gorm:"column:create_time;autoCreateTime"`
	UpdateTime   *time.Time `gorm:"column:update_time;autoUpdateTime"`
	LastLogin    *time.Time `gorm:"column:last_login"`
	IsDeleted    bool       `gorm:"column:is_deleted"` // 关键点：必须包含此字段
}

func (AdminUser) TableName() string {
	return "admin_user"
}
