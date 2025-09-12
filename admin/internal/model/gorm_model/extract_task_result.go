package gorm_model

import (
	"time"
)

type Entity struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	EntityName string `gorm:"column:entity_name; type:varchar(255); not null; comment:实体名称" json:"entity_name"`

	TaskID int `gorm:"column:task_id; type:int(11); not null; comment:任务ID" json:"task_id"`

	OntologyID int `gorm:"column:ontology_id; type:int(11); not null; comment:本体ID" json:"ontology_id"`
}

func (u *Entity) TableName() string {
	return "entity"
}

type Prop struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	EntityID  int `gorm:"column:entity_id; type:int(11); not null; comment:实体ID" json:"entity_id"`

	PropName  string `gorm:"column:prop_name; type:varchar(255); not null; comment:属性名称" json:"prop_name"`
	PropValue string `gorm:"column:prop_value; type:varchar(255); not null; comment:属性值" json:"prop_value"`

	TaskID int `gorm:"column:task_id; type:int(11); not null; comment:任务ID" json:"task_id"`
}

func (u *Prop) TableName() string {
	return "prop"
}

type Relationship struct {
	ID               uint `gorm:"primarykey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	SourceEntityID   int    `gorm:"column:source_entity_id; type:int(11); not null; comment:源实体ID" json:"source_entity_id"`
	TargetEntityID   int    `gorm:"column:target_entity_id; type:int(11); not null; comment:目标实体ID" json:"target_entity_id"`
	RelationshipName string `gorm:"column:relationship_name; type:varchar(255); not null; comment:关系名称" json:"relationship_name"`

	TaskID int `gorm:"column:task_id; type:int(11); not null; comment:任务ID" json:"task_id"`
}

func (u *Relationship) TableName() string {
	return "relationship"
}
