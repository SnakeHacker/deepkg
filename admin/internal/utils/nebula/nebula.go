package nebula

import (
	"time"

	"github.com/golang/glog"
	nebula "github.com/vesoft-inc/nebula-go/v3"
)

type NebulaSessionConfig struct {
	Timeout         int
	IdleTime        int
	MaxConnPoolSize int
	MinConnPoolSize int

	Host     string
	Port     int
	Username string
	Password string
}

func NewNebulaSession(conf NebulaSessionConfig) (session *nebula.Session, err error) {

	config := nebula.PoolConfig{
		TimeOut:         time.Duration(conf.Timeout) * time.Millisecond,
		IdleTime:        time.Duration(conf.IdleTime) * time.Millisecond,
		MaxConnPoolSize: conf.MaxConnPoolSize,
		MinConnPoolSize: conf.MinConnPoolSize,
	}

	pool, err := nebula.NewConnectionPool([]nebula.HostAddress{{Host: conf.Host, Port: conf.Port}}, config, nebula.DefaultLogger{})
	if err != nil {
		glog.Error(err)
		return
	}

	session, err = pool.GetSession(conf.Username, conf.Password)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
