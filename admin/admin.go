package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"

	"github.com/SnakeHacker/deepkg/admin/internal/config"
	"github.com/SnakeHacker/deepkg/admin/internal/handler"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/utils/io"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Set("log_dir", "./logs")
	flag.Set("alsologtostderr", "true")
	flag.Parse()

	err := io.CreateDirIfNotExist("./logs")
	if err != nil {
		glog.Fatal(err)
	}

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.MustSetup(c.Log)
	logx.Disable()

	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(func(header http.Header) {
		header.Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
		header.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Cache-Control")
		header.Set("Access-Control-Allow-Origin", "*")
	}, nil, "*"))

	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		close(idleConnsClosed)
	}()

	go func() {
		glog.Infof("Starting server at %s:%d...\n", c.Host, c.Port)
		server.Start()
	}()

	<-idleConnsClosed
	glog.Flush()
}
