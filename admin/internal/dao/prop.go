package dao

import (
	"errors"

	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"
	"gorm.io/gorm"
)

func CreateProp(db *gorm.DB, prop *m.Prop) (err error) {
	if prop == nil {
		err = errors.New("missing prop object")
		glog.Error(err)
		return
	}
	if err = db.Create(prop).Error; err != nil {
		glog.Error(err)
		return
	}

	return
}

func DeletePropsByIDs(db *gorm.DB, ids []int64) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&m.Prop{}).Error
	if err != nil {
		err = errors.New("prop is not existed")
		glog.Error(err)
		return
	}

	return
}

func DeletePropsByEntityIDs(db *gorm.DB, ids []int64) (err error) {
	err = db.Where("entity_id IN (?)", ids).Delete(&m.Prop{}).Error
	if err != nil {
		err = errors.New("prop is not existed")
		glog.Error(err)
		return
	}

	return
}

func SelectProps(db *gorm.DB, entityID int, pageIndex int, pageSize int) (props []*m.Prop, total int64, err error) {
	statement := db.Model(&m.Prop{})
	if entityID != 0 {
		statement = statement.Where("entity_id = ?", entityID)
	}

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

func SelectPropsByIDs(db *gorm.DB, entityIDs []int) (props []*m.Prop, err error) {
	statement := db.Model(&m.Prop{}).Where("entity_id IN (?) ", entityIDs)

	err = statement.Order("created_at desc").Distinct().Find(&props).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectPropByID(db *gorm.DB, id int) (prop m.Prop, err error) {
	err = db.Where("id = ?", id).First(&prop).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func UpdateProp(db *gorm.DB, prop *m.Prop) (err error) {
	if prop == nil {
		err = errors.New("missing prop object")
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
