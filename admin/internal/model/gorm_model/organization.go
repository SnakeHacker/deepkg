package gorm_model

import (
	"time"
)

type Organization struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	OrgName   string `gorm:"column:org_name; type:varchar(255); not null; comment:组织名称" json:"org_name"`
}

func (o *Organization) TableName() string {
	return "organization"
}
