package dao

import (
	"errors"

	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"
	"gorm.io/gorm"
)

func CreateSchemaOntologyProp(db *gorm.DB, prop *m.SchemaOntologyProp) (err error) {
	if prop == nil {
		err = errors.New("missing schema ontology prop object")
		glog.Error(err)
		return
	}
	if err = db.Create(prop).Error; err != nil {
		glog.Error(err)
		return
	}

	return
}

func DeleteSchemaOntologyPropsByIDs(db *gorm.DB, ids []int64) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&m.SchemaOntologyProp{}).Error
	if err != nil {
		err = errors.New("schema ontology prop is not existed")
		glog.Error(err)
		return
	}

	return
}

func SelectSchemaOntologyProps(db *gorm.DB, ontologyID int, pageIndex int, pageSize int) (props []*m.SchemaOntologyProp, total int64, err error) {

	statement := db.Model(&m.SchemaOntologyProp{}).Where("ontology_id = ?", ontologyID)

	err = statement.Count(&total).Error
	if err != nil {
		glog.Error(err)
		return
	}

	if pageIndex != -1 {
		statement = statement.Offset((pageIndex - 1) * pageSize).Limit(pageSize)
	}

	err = statement.Order("created_at desc").Distinct().Find(&props).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectSchemaOntologyPropByID(db *gorm.DB, id int64) (prop m.SchemaOntologyProp, err error) {
	err = db.Where("id = ?", id).First(&prop).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func UpdateSchemaOntologyProp(db *gorm.DB, prop *m.SchemaOntologyProp) (err error) {
	if prop == nil {
		err = errors.New("missing schema ontology prop object")
		glog.Error(err)
		return
	}

	err = db.Save(prop).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
