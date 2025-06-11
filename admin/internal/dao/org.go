package dao

import (
	"errors"

	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/golang/glog"
	"gorm.io/gorm"
)

func CreateOrg(db *gorm.DB, org *m.Organization) (err error) {
	if org == nil {
		err = errors.New("organization object is missing")
		glog.Error(err)
		return
	}
	if err = db.Create(org).Error; err != nil {
		glog.Error(err)
		return
	}

	return
}

func DeleteOrgsByIDs(db *gorm.DB, ids []int64) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&m.Organization{}).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectOrgs(db *gorm.DB, pageIndex int, pageSize int) (orgs []*m.Organization, total int64, err error) {
	statement := db.Model(&m.Organization{})

	err = statement.Count(&total).Error
	if err != nil {
		glog.Error(err)
		return
	}

	if pageIndex != -1 {
		statement = statement.Offset((pageIndex - 1) * pageSize).Limit(pageSize)
	}

	err = statement.Order("created_at desc").Distinct().Find(&orgs).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectOrgByID(db *gorm.DB, id int64) (org *m.Organization, err error) {
	err = db.Where("id = ?", id).First(&org).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func UpdateOrg(db *gorm.DB, organization *m.Organization) (err error) {
	if organization == nil {
		err = errors.New("organization object is missing")
		glog.Error(err)
		return
	}

	err = db.Save(organization).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func SelectUsersByOrgIDs(db *gorm.DB, ids []int64) (users []*m.User, err error) {
	err = db.Where("org_id IN (?)", ids).Find(&users).Error
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func CheckOrgNameExists(db *gorm.DB, orgName string) (exists bool, err error) {
	var org *m.Organization
	err = db.Where("org_name = ?", orgName).First(&org).Error
	if err != nil {
		glog.Error(err)
		return false, err // 查询出错
	}
	return true, nil // 如果查询到组织，则返回存在
}
