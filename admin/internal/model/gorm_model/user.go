package gorm_model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserCode     string `gorm:"column:user_code; type:varchar(255); unique; not null; comment:用户编码" json:"user_code"`
	OrgID        int    `gorm:"column:org_id; type:int(11); not null; default:0; comment:组织ID" json:"org_id"`
	Account      string `gorm:"column:account; type:varchar(255); unique; not null; comment:账号用于登录" json:"account"`
	Username     string `gorm:"column:username; type:varchar(255); not null; comment:用户名称" json:"username"`
	PasswordHash string `gorm:"column:password_hash; type:varchar(255); not null; comment:密码" json:"password_hash"`
	Phone        string `gorm:"column:phone; type:varchar(255); not null; default:''; comment:电话" json:"phone"`
	Mail         string `gorm:"column:mail; type:varchar(255); not null; default:''; comment:邮箱" json:"mail"`
	Enable       int    `gorm:"column:enable; type:tinyint(1); not null; comment:启用状态：1-启用，2-禁用" json:"enable"`
	Role         int    `gorm:"column:role; type:tinyint(1); not null; default:0; comment:角色：2-普通用户，1-管理员" json:"role"`
	Avatar       string `gorm:"column:avatar; type:varchar(255); not null; default:''; comment:头像" json:"avatar"`
}

func (u *User) TableName() string {
	return "user"
}
