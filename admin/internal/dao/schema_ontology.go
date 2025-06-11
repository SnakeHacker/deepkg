package dao

import (
	"errors"

	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"
	"gorm.io/gorm"
)

func CreateSchemaOntology(db *gorm.DB, ontology *m.SchemaOntology) (err error) {
	if ontology == nil {
		err = errors.New("missing schema ontology object")
		glog.Error(err)
		return
	}
	if err = db.Create(ontology).Error; err != nil {
		glog.Error(err)
		return
	}

	return
}

func DeleteSchemaOntologysByIDs(db *gorm.DB, ids []int64) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&m.SchemaOntology{}).Error
	if err != nil {
		err = errors.New("schema ontology is not existed")
		glog.Error(err)
		return
	}

	return
}

func SelectSchemaOntologys(db *gorm.DB, workspaceID int) (ontologys []*m.SchemaOntology, total int64, err error) {

	statement := db.Model(&m.SchemaOntology{}).Where("work_space_id = ?", workspaceID)

	err = statement.Count(&total).Error
	if err != nil {
		glog.Error(err)
		return
	}

	err = statement.Order("created_at desc").Distinct().Find(&ontologys).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectSchemaOntologyByID(db *gorm.DB, id int64) (ontology m.SchemaOntology, err error) {
	err = db.Where("id = ?", id).First(&ontology).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func UpdateSchemaOntology(db *gorm.DB, ontology *m.SchemaOntology) (err error) {
	if ontology == nil {
		err = errors.New("missing schema ontology object")
		glog.Error(err)
		return
	}

	err = db.Save(ontology).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
