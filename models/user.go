package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Username string `gorm:"column:username" json:"username"` // 用户名
	Password string `gorm:"column:password" json:"-"`        // 密码
	Mobile   string `gorm:"column:mobile" json:"mobile"`     // 手机号
	Email    string `gorm:"column:email" json:"email"`       // 邮箱
	Role     int    `gorm:"column:role" json:"role"`         // 角色

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (*User) TableName() string {
	return "user"
}
