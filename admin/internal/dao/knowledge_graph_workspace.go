package dao

import (
	"errors"

	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"
	"gorm.io/gorm"
)

func CreateKnowledgeGraphWorkspace(db *gorm.DB, wsp *m.KnowledgeGraphWorkspace) (err error) {
	if wsp == nil {
		err = errors.New("missing knowledge graph workspace object")
		glog.Error(err)
		return
	}
	if err = db.Create(wsp).Error; err != nil {
		glog.Error(err)
		return
	}

	return
}

func DeleteKnowledgeGraphWorkspacesByIDs(db *gorm.DB, ids []int64) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&m.KnowledgeGraphWorkspace{}).Error
	if err != nil {
		err = errors.New("knowledge graph workspace is not existed")
		glog.Error(err)
		return
	}

	return
}

func SelectKnowledgeGraphWorkspaces(db *gorm.DB, pageIndex int, pageSize int) (wsps []*m.KnowledgeGraphWorkspace, total int64, err error) {

	statement := db.Model(&m.KnowledgeGraphWorkspace{})

	err = statement.Count(&total).Error
	if err != nil {
		glog.Error(err)
		return
	}

	if pageIndex != -1 {
		statement = statement.Offset((pageIndex - 1) * pageSize).Limit(pageSize)
	}

	err = statement.Order("created_at desc").Distinct().Find(&wsps).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectKnowledgeGraphWorkspaceByID(db *gorm.DB, id int64) (wsp m.KnowledgeGraphWorkspace, err error) {
	err = db.Where("id = ?", id).First(&wsp).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func UpdateKnowledgeGraphWorkspace(db *gorm.DB, wsp *m.KnowledgeGraphWorkspace) (err error) {
	if wsp == nil {
		err = errors.New("missing knowledge graph workspace object")
		glog.Error(err)
		return
	}

	err = db.Save(wsp).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
