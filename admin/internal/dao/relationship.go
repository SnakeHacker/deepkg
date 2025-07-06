package dao

import (
	"errors"

	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"
	"gorm.io/gorm"
)

func CreateRelationship(db *gorm.DB, relationship *m.Relationship) (err error) {
	if relationship == nil {
		err = errors.New("missing relationship object")
		glog.Error(err)
		return
	}
	if err = db.Create(relationship).Error; err != nil {
		glog.Error(err)
		return
	}

	return
}

func DeleteRelationshipsByIDs(db *gorm.DB, ids []int64) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&m.Relationship{}).Error
	if err != nil {
		err = errors.New("relationship is not existed")
		glog.Error(err)
		return
	}

	return
}

func DeleteRelationshipsByEntityIDs(db *gorm.DB, ids []int64) (err error) {
	err = db.Where("source_entity_id IN (?) OR target_entity_id IN (?)", ids, ids).Delete(&m.Relationship{}).Error
	if err != nil {
		err = errors.New("relationship is not existed")
		glog.Error(err)
		return
	}

	return
}

func SelectRelationshipsByTaskID(db *gorm.DB, taskID int, pageIndex int, pageSize int) (relationships []*m.Relationship, total int64, err error) {
	statement := db.Model(&m.Relationship{})
	if taskID != 0 {
		statement = statement.Where("task_id = ?", taskID)
	}

	err = statement.Count(&total).Error
	if err != nil {
		glog.Error(err)
		return
	}

	if pageIndex != -1 {
		statement = statement.Offset((pageIndex - 1) * pageSize).Limit(pageSize)
	}

	err = statement.Order("created_at desc").Distinct().Find(&relationships).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectRelationshipByID(db *gorm.DB, id int) (relationship m.Relationship, err error) {
	err = db.Where("id = ?", id).First(&relationship).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func UpdateRelationship(db *gorm.DB, relationship *m.Relationship) (err error) {
	if relationship == nil {
		err = errors.New("missing relationship object")
		glog.Error(err)
		return
	}

	err = db.Save(relationship).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
