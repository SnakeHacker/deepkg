package dao

import (
	"errors"

	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"
	"gorm.io/gorm"
)

func CreateExtractTask(db *gorm.DB, task *m.ExtractTask) (err error) {
	if task == nil {
		err = errors.New("missing task object")
		glog.Error(err)
		return
	}
	if err = db.Create(task).Error; err != nil {
		glog.Error(err)
		return
	}

	return
}

func DeleteExtractTasksByIDs(db *gorm.DB, ids []int64) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&m.ExtractTask{}).Error
	if err != nil {
		err = errors.New("task is not existed")
		glog.Error(err)
		return
	}

	return
}

func SelectExtractTasks(db *gorm.DB, workspaceID int, pageIndex int, pageSize int) (tasks []*m.ExtractTask, total int64, err error) {
	statement := db.Model(&m.ExtractTask{})
	if workspaceID != 0 {
		statement = statement.Where("work_space_id = ?", workspaceID)
	}

	err = statement.Count(&total).Error
	if err != nil {
		glog.Error(err)
		return
	}

	if pageIndex != -1 {
		statement = statement.Offset((pageIndex - 1) * pageSize).Limit(pageSize)
	}

	err = statement.Order("created_at desc").Distinct().Find(&tasks).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectExtractTaskByID(db *gorm.DB, id int) (task m.ExtractTask, err error) {
	err = db.Where("id = ?", id).First(&task).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func UpdateExtractTask(db *gorm.DB, task *m.ExtractTask) (err error) {
	if task == nil {
		err = errors.New("missing task object")
		glog.Error(err)
		return
	}

	err = db.Save(task).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func UpdateExtractTaskStatus(db *gorm.DB, taskID int, taskStatus int) (err error) {
	err = db.Model(&m.ExtractTask{}).Where("id = ?", taskID).Update("task_status", taskStatus).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func UpdateExtractTaskPublished(db *gorm.DB, taskID int, published bool) (err error) {
	err = db.Model(&m.ExtractTask{}).Where("id = ?", taskID).Update("published", published).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
