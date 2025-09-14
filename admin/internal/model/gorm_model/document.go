package gorm_model

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	DocName   string `gorm:"column:doc_name; type:varchar(255); not null; comment:文件名" json:"doc_name"`
	DocDesc   string `gorm:"column:doc_desc; type:text; not null; comment:文件描述" json:"doc_desc"`
	DocPath   string `gorm:"column:doc_path; type:varchar(255); not null; comment:文件路径" json:"doc_path"`
	DirID     int    `gorm:"column:dir_id; type:int(11); not null; comment:文件目录ID" json:"dir_id"`
	CreatorID int    `gorm:"column:creator_id; type:int(11); not null; comment:创建者ID" json:"creator_id"`
}

func (u *Document) TableName() string {
	return "document"
}
