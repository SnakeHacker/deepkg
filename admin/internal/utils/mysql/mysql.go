package mysql

import (
	"time"

	"github.com/SnakeHacker/deepkg/admin/internal/config"
	"github.com/golang/glog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(conf config.Config) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(conf.Mysql.Datasource), &gorm.Config{})
	if err != nil {
		glog.Error(err)
		return
	}

	dbConn, err := db.DB()
	if err != nil {
		glog.Fatal(err)
		return
	}

	dbConn.SetMaxIdleConns(10)
	dbConn.SetMaxOpenConns(100)
	dbConn.SetConnMaxLifetime(60 * time.Second)

	return
}
