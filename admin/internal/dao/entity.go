package dao

import (
	"errors"

	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"
	"gorm.io/gorm"
)

func CreateEntity(db *gorm.DB, entity *m.Entity) (err error) {
	if entity == nil {
		err = errors.New("missing entity object")
		glog.Error(err)
		return
	}
	if err = db.Create(entity).Error; err != nil {
		glog.Error(err)
		return
	}

	return
}

func DeleteEntitysByIDs(db *gorm.DB, ids []int64) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&m.Entity{}).Error
	if err != nil {
		err = errors.New("entity is not existed")
		glog.Error(err)
		return
	}

	return
}

func SelectEntitiesByTaskID(db *gorm.DB, taskID int, pageIndex int, pageSize int) (entities []*m.Entity, total int64, err error) {
	statement := db.Model(&m.Entity{})
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

	err = statement.Order("created_at desc").Distinct().Find(&entities).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectEntityByID(db *gorm.DB, id int) (entity m.Entity, err error) {
	err = db.Where("id = ?", id).First(&entity).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectEntityModelsByIDs(db *gorm.DB, ids []int64) (entities []*m.Entity, err error) {
	err = db.Where("id IN (?)", ids).Find(&entities).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func UpdateEntity(db *gorm.DB, entity *m.Entity) (err error) {
	if entity == nil {
		err = errors.New("missing entity object")
		glog.Error(err)
		return
	}

	err = db.Save(entity).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
