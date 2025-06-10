package dao

import (
	"errors"

	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"
	"gorm.io/gorm"
)

func CreateDocumentDir(db *gorm.DB, dir *m.DocumentDir) (err error) {
	if dir == nil {
		err = errors.New("missing document dir object")
		glog.Error(err)
		return
	}
	if err = db.Create(dir).Error; err != nil {
		glog.Error(err)
		return
	}

	return
}

func DeleteDocumentDirsByIDs(db *gorm.DB, ids []int64) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&m.DocumentDir{}).Error
	if err != nil {
		err = errors.New("document dir is not existed")
		glog.Error(err)
		return
	}

	return
}

func SelectDocumentDirsByParentIDs(db *gorm.DB, parentIDs []int64) (dirs []*m.DocumentDir, total int64, err error) {

	statement := db.Model(&m.DocumentDir{})
	if len(parentIDs) > 0 {
		statement = statement.Where("parent_id IN (?)", parentIDs)
	}

	err = statement.Count(&total).Error
	if err != nil {
		glog.Error(err)
		return
	}

	err = statement.Order("created_at desc").Distinct().Find(&dirs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		glog.Error(err)
		return
	}

	return
}

func SelectDocumentDirs(db *gorm.DB) (dirs []*m.DocumentDir, total int64, err error) {

	statement := db.Model(&m.DocumentDir{})

	err = statement.Count(&total).Error
	if err != nil {
		glog.Error(err)
		return
	}

	err = statement.Order("created_at desc").Distinct().Find(&dirs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		glog.Error(err)
		return
	}

	return
}

func SelectDocumentDirByID(db *gorm.DB, id int64) (dir m.DocumentDir, err error) {
	err = db.Where("id = ?", id).First(&dir).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.New("dir is not existed")
		}
		glog.Error(err)
		return
	}

	return
}

func UpdateDocumentDir(db *gorm.DB, dir *m.DocumentDir) (err error) {
	if dir == nil {
		err = errors.New("missing document dir object")
		glog.Error(err)
		return
	}

	err = db.Save(dir).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectDocumentDirByName(db *gorm.DB, dirName string) (dir m.DocumentDir, err error) {
	err = db.Where("dir_name = ?", dirName).First(&dir).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
