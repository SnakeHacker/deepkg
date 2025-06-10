package main

import (
	"flag"

	"github.com/SnakeHacker/deepkg/admin/internal/config"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/utils/mysql"
	"github.com/zeromicro/go-zero/core/conf"

	"github.com/golang/glog"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	db, err := mysql.NewMySQL(c)
	if err != nil {
		glog.Error(err)
		return
	}

	err = mysql.ResetTables(db,
		&gorm_model.User{},
		&gorm_model.Organization{},
		&gorm_model.DocumentDir{},
		&gorm_model.Document{},
		&gorm_model.KnowledgeGraphWorkspace{},
		&gorm_model.SchemaOntology{},
		&gorm_model.SchemaOntologyProp{},
		&gorm_model.SchemaTriple{},
	)

	if err != nil {
		glog.Error(err)
		return
	}
}
