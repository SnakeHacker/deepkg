package gorm_model

import (
	"gorm.io/gorm"
)

type SchemaOntology struct {
	gorm.Model
	WorkSpaceID  int    `gorm:"column:work_space_id; type:int(11); not null; comment:工作空间ID" json:"work_space_id"`
	OntologyName string `gorm:"column:ontology_name; type:varchar(255); not null; comment:实体名称" json:"ontology_name"`
	OntologyDesc string `gorm:"column:ontology_desc; type:text; not null; comment:实体描述" json:"ontology_desc"`

	CreatorID int `gorm:"column:creator_id; type:int(11); not null; comment:创建者ID" json:"creator_id"`
}

func (u *SchemaOntology) TableName() string {
	return "schema_ontology"
}

type SchemaOntologyProp struct {
	gorm.Model
	WorkSpaceID int    `gorm:"column:work_space_id; type:int(11); not null; comment:工作空间ID" json:"work_space_id"`
	OntologyID  int    `gorm:"column:ontology_id; type:int(11); not null; comment:实体ID" json:"ontology_id"`
	PropName    string `gorm:"column:prop_name; type:varchar(255); not null; comment:属性名称" json:"prop_name"`
	PropDesc    string `gorm:"column:prop_desc; type:text; not null; comment:属性描述" json:"prop_desc"`

	CreatorID int `gorm:"column:creator_id; type:int(11); not null; comment:创建者ID" json:"creator_id"`
}

func (u *SchemaOntologyProp) TableName() string {
	return "schema_ontology_prop"
}

// 三元组
type SchemaTriple struct {
	gorm.Model
	WorkSpaceID      int    `gorm:"column:work_space_id; type:int(11); not null; comment:工作空间ID" json:"work_space_id"`
	SourceOntologyID int    `gorm:"column:source_ontology_id; type:int(11); not null; comment:源实体ID" json:"source_ontology_id"`
	TargetOntologyID int    `gorm:"column:target_ontology_id; type:int(11); not null; comment:目标实体ID" json:"target_ontology_id"`
	Relationship     string `gorm:"column:relationship; type:varchar(255); not null; comment:实体关系" json:"relationship"`

	CreatorID int `gorm:"column:creator_id; type:int(11); not null; comment:创建者ID" json:"creator_id"`
}

func (u *SchemaTriple) TableName() string {
	return "schema_triple"
}
