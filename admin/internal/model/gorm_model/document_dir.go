package gorm_model

import (
	"gorm.io/gorm"
)

type DocumentDir struct {
	gorm.Model
	DirName   string `gorm:"column:dir_name; type:varchar(255); not null; comment:目录名" json:"dir_name"`
	ParentID  int    `gorm:"column:parent_id; type:int(11); not null; default:0; comment:父目录ID" json:"parent_id"`
	SortIndex int    `gorm:"column:sort_index; type:int(11); not null; default:0; comment:排序索引" json:"sort_index"`
	Remark    string `gorm:"column:remark; type:text; not null; default:''; comment:备注" json:"remark"`
}

func (u *DocumentDir) TableName() string {
	return "DocumentDir"
}
