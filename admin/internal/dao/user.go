package dao

import (
	"errors"

	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *m.User) (err error) {
	if user == nil {
		err = errors.New("缺少 user用户 对象")
		glog.Error(err)
		return
	}
	if err = db.Create(user).Error; err != nil {
		glog.Error(err)
		return
	}

	return
}

func DeleteUsersByIDs(db *gorm.DB, ids []int64) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&m.User{}).Error
	if err != nil {
		err = errors.New("User用户对象不存在")
		glog.Error(err)
		return
	}

	return
}

func SelectUsers(db *gorm.DB, pageIndex int, pageSize int) (users []*types.User, total int64, err error) {

	statement := db.Table("user").Select("user.*, organization.org_name").
		Joins("JOIN organization ON user.org_id = organization.id").Where("user.deleted_at IS NULL")

	err = statement.Count(&total).Error
	if err != nil {
		glog.Error(err)
		return
	}

	if pageIndex != -1 {
		statement = statement.Offset((pageIndex - 1) * pageSize).Limit(pageSize)
	}

	err = statement.Order("created_at desc").Distinct().Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		glog.Error(err)
		return
	}

	return
}

func SelectUserByID(db *gorm.DB, id int64) (user *types.User, err error) {
	user = &types.User{}
	statement := db.Table("user").Select("user.*, organization.org_name").
		Joins("JOIN organization ON user.org_id = organization.id").
		Where("user.id = ?", id)

	err = statement.First(user).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectUserModelByID(db *gorm.DB, id int64) (user *m.User, err error) {
	err = db.Where("id = ?", id).First(&user).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func UpdateUser(db *gorm.DB, user *m.User) (err error) {
	if user == nil {
		err = errors.New("缺少User用户对象")
		glog.Error(err)
		return
	}

	err = db.Save(user).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectUserByAccount(db *gorm.DB, account string) (user m.User, err error) {
	err = db.Where("account = ?", account).First(&user).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectUserModelsByIDs(db *gorm.DB, ids []int64) (users []*m.User, err error) {
	err = db.Where("id IN (?)", ids).Find(&users).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
