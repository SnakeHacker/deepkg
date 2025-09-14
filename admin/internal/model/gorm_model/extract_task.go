package gorm_model

import (
	"gorm.io/gorm"
)

type ExtractTask struct {
	gorm.Model
	TaskName    string `gorm:"column:task_name; type:varchar(255); not null; comment:任务名称" json:"task_name"`
	Remark      string `gorm:"column:remark; type:text; not null; comment:备注" json:"remark"`
	WorkSpaceID int64  `gorm:"column:work_space_id; type:int(11); not null; comment:工作空间ID" json:"work_space_id"`

	TaskStatus int  `gorm:"column:task_status; type:int(11); not null; comment:任务状态" json:"task_status"`
	Published  bool `gorm:"column:published; type:tinyint(1); not null; comment:是否发布" json:"published"`

	CreatorID int `gorm:"column:creator_id; type:int(11); not null; comment:创建者ID" json:"creator_id"`
}

func (u *ExtractTask) TableName() string {
	return "extract_task"
}

type ExtractTaskDocument struct {
	gorm.Model
	TaskID int `gorm:"column:task_id; type:int(11); not null; comment:任务ID" json:"task_id"`
	DocID  int `gorm:"column:doc_id; type:int(11); not null; comment:文档ID" json:"doc_id"`
}

func (u *ExtractTaskDocument) TableName() string {
	return "extract_task_document"
}

type ExtractTaskTriple struct {
	gorm.Model
	TaskID   int `gorm:"column:task_id; type:int(11); not null; comment:任务ID" json:"task_id"`
	TripleID int `gorm:"column:triple_id; type:int(11); not null; comment:三元组ID" json:"triple_id"`
}

func (u *ExtractTaskTriple) TableName() string {
	return "extract_task_triple"
}
