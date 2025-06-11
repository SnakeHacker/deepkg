package config

import (
	"github.com/SnakeHacker/deepkg/admin/internal/utils/nebula"
	"github.com/SnakeHacker/deepkg/admin/internal/utils/s3/minio"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Minio minio.MinioConf

	Auth struct {
		AccessSecret     string
		AccessExpire     int64
		HardAccessSecret string
	}

	Mysql struct {
		Datasource string
	}

	Redis struct {
		IsCluster bool
		Hosts     []string
		Pass      string
	}

	Nebula nebula.NebulaSessionConfig
}
