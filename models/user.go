package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UUID     string `gorm:"column:uuid;type:varchar(32);" json:"uuid"`         // 用户的唯一标识
	UserName string `gorm:"column:username;type:varchar(64);" json:"username"` // 用户名
	Password string `gorm:"column:password;type:varchar(64);" json:"password"` // 密码
	Mobile   string `gorm:"column:mobile;type:varchar(32);" json:"mobile"`     // 手机号
	Email    string `gorm:"column:email;type:varchar(32);" json:"email"`       // 邮箱
	Role     int    `gorm:"column:is_admin;type:int;" json:"role"`             // 是否是管理员[0-否, 1-是]
}

func (table *User) TableName() string {
	return "user"
}
