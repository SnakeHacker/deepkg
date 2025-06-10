package svc

import (
	"crypto/rsa"

	"github.com/SnakeHacker/deepkg/admin/common/captcha"
	rsa2 "github.com/SnakeHacker/deepkg/admin/common/rsa"
	"github.com/SnakeHacker/deepkg/admin/internal/config"
	"github.com/SnakeHacker/deepkg/admin/internal/middleware"
	"github.com/SnakeHacker/deepkg/admin/internal/utils/mysql"
	"github.com/SnakeHacker/deepkg/admin/internal/utils/s3/minio"
	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
	"github.com/golang/glog"
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config     config.Config
	DB         *gorm.DB
	JwtX       rest.Middleware
	HTTPClient *resty.Client
	Minio      *minio.Client
	Redis      redis.UniversalClient
	PrivateKey *rsa.PrivateKey
	Captcha    *base64Captcha.Captcha
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := mysql.NewMySQL(c)
	if err != nil {
		panic(err)
	}

	httpClient := resty.New()

	// Init minio client
	minioClient, err := minio.NewMinioClient(c.Minio)
	if err != nil {
		glog.Fatal(err)
	}

	// Init redis client
	redisClient := NewRedisClient(c)

	// Init RSA key
	privateKey, err := rsa2.GenerateKey(2048)
	if err != nil {
		glog.Fatal(err)
	}

	// Init Captcha
	svcCaptcha, err := captcha.SetUpCaptcha(nil)
	if err != nil {
		glog.Fatal(err)
	}

	return &ServiceContext{
		Config:     c,
		DB:         db,
		HTTPClient: httpClient,
		Minio:      minioClient,
		Redis:      redisClient,
		PrivateKey: privateKey,
		JwtX:       middleware.NewJwtXMiddleware(redisClient, c).Handle,
		Captcha:    svcCaptcha,
	}
}

func NewRedisClient(c config.Config) redis.UniversalClient {
	if c.Redis.IsCluster {
		return redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    c.Redis.Hosts,
			Password: c.Redis.Pass,
		})
	}

	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Hosts[0],
		Password: c.Redis.Pass,
	})
}
