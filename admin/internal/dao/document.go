package dao

import (
	"errors"

	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"
	"gorm.io/gorm"
)

func CreateDocument(db *gorm.DB, doc *m.Document) (err error) {
	if doc == nil {
		err = errors.New("missing document object")
		glog.Error(err)
		return
	}
	if err = db.Create(doc).Error; err != nil {
		glog.Error(err)
		return
	}

	return
}

func DeleteDocumentsByIDs(db *gorm.DB, ids []int64) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&m.Document{}).Error
	if err != nil {
		err = errors.New("document is not existed")
		glog.Error(err)
		return
	}

	return
}

func SelectDocuments(db *gorm.DB, dirID int64, pageIndex int, pageSize int) (docs []types.Document, total int64, err error) {

	statement := db.Model(&m.Document{})

	statement = db.Table("document").Select("document.*, user.id, user.username").
		Joins("JOIN user ON document.creator_id = user.id").Where("user.deleted_at IS NULL")

	if dirID > 0 {
		statement = statement.Where("document.dir_id = ?", dirID)
	}

	err = statement.Count(&total).Error
	if err != nil {
		glog.Error(err)
		return
	}

	if pageIndex != -1 {
		statement = statement.Offset((pageIndex - 1) * pageSize).Limit(pageSize)
	}

	err = statement.Order("created_at desc").Distinct().Find(&docs).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectDocumentByID(db *gorm.DB, id int64) (doc types.Document, err error) {
	statement := db.Table("document").Select("document.*, user.id, user.username").
		Joins("JOIN user ON document.creator_id = user.id").Where("user.deleted_at IS NULL AND document.id = ?")

	err = statement.First(&doc).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectDocumentModelByID(db *gorm.DB, id int64) (doc *m.Document, err error) {
	err = db.Where("id = ?", id).First(&doc).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func UpdateDocument(db *gorm.DB, doc *m.Document) (err error) {
	if doc == nil {
		err = errors.New("missing document object")
		glog.Error(err)
		return
	}

	err = db.Save(doc).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
