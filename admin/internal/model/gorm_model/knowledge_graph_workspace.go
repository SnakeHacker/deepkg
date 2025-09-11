package gorm_model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type KnowledgeGraphWorkspace struct {
	ID            uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     soft_delete.DeletedAt `gorm:"uniqueIndex:udx_work_space_name;"`
	WorkSpaceName string                `gorm:"column:work_space_name; type:varchar(255); not null; comment:知识库名称; uniqueIndex:udx_work_space_name" json:"work_space_name"`
	CreatorID     int                   `gorm:"column:creator_id; type:int(11); not null; comment:创建者ID" json:"creator_id"`
}

func (u *KnowledgeGraphWorkspace) TableName() string {
	return "knowledge_graph_workspace"
}
