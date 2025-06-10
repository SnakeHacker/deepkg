package gorm_model

import (
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	OrgName string `gorm:"column:org_name; type:varchar(255); not null; comment:组织名称" json:"org_name"`
}

func (o *Organization) TableName() string {
	return "organization"
}
