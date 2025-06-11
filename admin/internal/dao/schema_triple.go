package dao

import (
	"errors"

	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"
	"gorm.io/gorm"
)

func CreateSchemaTriple(db *gorm.DB, triple *m.SchemaTriple) (err error) {
	if triple == nil {
		err = errors.New("missing schema triple object")
		glog.Error(err)
		return
	}
	if err = db.Create(triple).Error; err != nil {
		glog.Error(err)
		return
	}

	return
}

func DeleteSchemaTriplesByIDs(db *gorm.DB, ids []int64) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&m.SchemaTriple{}).Error
	if err != nil {
		err = errors.New("schema triple is not existed")
		glog.Error(err)
		return
	}

	return
}

func SelectSchemaTriples(db *gorm.DB, workspaceID int, pageIndex int, pageSize int) (triples []*m.SchemaTriple, total int64, err error) {

	statement := db.Model(&m.SchemaTriple{}).Where("work_space_id = ?", workspaceID)

	err = statement.Count(&total).Error
	if err != nil {
		glog.Error(err)
		return
	}

	if pageIndex != -1 {
		statement = statement.Offset((pageIndex - 1) * pageSize).Limit(pageSize)
	}

	err = statement.Order("created_at desc").Distinct().Find(&triples).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectSchemaTripleByID(db *gorm.DB, id int64) (triples m.SchemaTriple, err error) {
	err = db.Where("id = ?", id).First(&triples).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func UpdateSchemaTriple(db *gorm.DB, triples *m.SchemaTriple) (err error) {
	if triples == nil {
		err = errors.New("missing schema triples object")
		glog.Error(err)
		return
	}

	err = db.Save(triples).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
