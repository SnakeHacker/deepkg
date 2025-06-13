package dao

import (
	"errors"

	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"
	"gorm.io/gorm"
)

func CreateExtractTaskDocument(db *gorm.DB, doc *m.ExtractTaskDocument) (err error) {
	if doc == nil {
		err = errors.New("missing task doc object")
		glog.Error(err)
		return
	}
	if err = db.Create(doc).Error; err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectExtractTaskDocuments(db *gorm.DB, taskID int) (docs []*m.ExtractTaskDocument, err error) {
	statement := db.Model(&m.ExtractTaskDocument{})
	if taskID == 0 {
		err = errors.New("taskID is nil")
		glog.Error(err)
		return
	}

	statement = statement.Where("task_id = ?", taskID)

	err = statement.Order("created_at desc").Distinct().Find(&docs).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
