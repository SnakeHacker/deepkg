package dao

import (
	"errors"

	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"
	"gorm.io/gorm"
)

func CreateExtractTaskTriple(db *gorm.DB, triple *m.ExtractTaskTriple) (err error) {
	if triple == nil {
		err = errors.New("missing task triple object")
		glog.Error(err)
		return
	}
	if err = db.Create(triple).Error; err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectExtractTaskTriples(db *gorm.DB, taskID int) (triples []*m.ExtractTaskTriple, err error) {
	statement := db.Model(&m.ExtractTaskTriple{})
	if taskID == 0 {
		err = errors.New("taskID is nil")
		glog.Error(err)
		return
	}

	statement = statement.Where("task_id = ?", taskID)

	err = statement.Order("created_at desc").Distinct().Find(&triples).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
