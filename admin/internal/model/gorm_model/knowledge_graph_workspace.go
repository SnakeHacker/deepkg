package gorm_model

import (
	"time"
)

type KnowledgeGraphWorkspace struct {
	ID            uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	WorkSpaceName string `gorm:"column:work_space_name; type:varchar(255); unique; not null; comment:知识库名称" json:"work_space_name"`

	CreatorID int `gorm:"column:creator_id; type:int(11); not null; comment:创建者ID" json:"creator_id"`
}

func (u *KnowledgeGraphWorkspace) TableName() string {
	return "knowledge_graph_workspace"
}
