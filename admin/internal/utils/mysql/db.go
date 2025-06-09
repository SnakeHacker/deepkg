package mysql

import (
	"fmt"

	"github.com/golang/glog"
	"gorm.io/gorm"
)

// CreateTables create all tables
func CreateTables(db *gorm.DB, values ...interface{}) (err error) {
	err = db.Migrator().CreateTable(values...)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

// DropTables drop all tables
func DropTables(db *gorm.DB, values ...interface{}) (err error) {
	err = db.Migrator().DropTable(values...)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

// ClearTables clear tabel data
func ClearTables(db *gorm.DB, values ...interface{}) (err error) {
	for _, value := range values {
		err = db.Exec(fmt.Sprintf(`DELETE FROM %s`, value)).Error
		if err != nil {
			glog.Error(err)
			return
		}
	}

	return
}

// ResetTables drop and create tables
func ResetTables(db *gorm.DB, values ...interface{}) (err error) {
	glog.Info("Droping tables")
	if err = DropTables(db, values...); err != nil {
		glog.Error(err)
		return
	}

	glog.Info("Creating tables")
	if err = CreateTables(db, values...); err != nil {
		glog.Error(err)
		return
	}

	return
}
