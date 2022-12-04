package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint           `gorm:"primarykey;" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
	UUID      string         `gorm:"column:uuid;type:varchar(50);" json:"uuid"`         // 用户的唯一标识
	UserName  string         `gorm:"column:username;type:varchar(50);" json:"username"` // 用户名
	Password  string         `gorm:"column:password;type:varchar(50);" json:"password"` // 密码
	Mobile    string         `gorm:"column:mobile;type:varchar(20);" json:"mobile"`     // 手机号
	Email     string         `gorm:"column:email;type:varchar(20);" json:"email"`       // 邮箱
	Role      int            `gorm:"column:is_admin;type:int;" json:"role"`             // 是否是管理员【0-否，1-是】
}

func (table *User) TableName() string {
	return "user"
}
