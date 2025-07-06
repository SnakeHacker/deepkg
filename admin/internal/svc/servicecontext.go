package svc

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"github.com/SnakeHacker/deepkg/admin/internal/utils/nebula"
	"net/http"

	"github.com/zeromicro/x/errors"

	"github.com/SnakeHacker/deepkg/admin/common/ai/llm"
	"github.com/SnakeHacker/deepkg/admin/common/captcha"
	rsa2 "github.com/SnakeHacker/deepkg/admin/common/rsa"
	"github.com/SnakeHacker/deepkg/admin/internal/config"
	"github.com/SnakeHacker/deepkg/admin/internal/middleware"
	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/utils/mysql"
	"github.com/SnakeHacker/deepkg/admin/internal/utils/s3/minio"
	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
	"github.com/golang/glog"
	"github.com/mojocn/base64Captcha"
	nebula_go "github.com/vesoft-inc/nebula-go/v3"
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
	Nebula     *nebula_go.Session
	PrivateKey *rsa.PrivateKey
	Captcha    *base64Captcha.Captcha
	LLM        llm.Client
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

	nebulaSession, err := nebula.NewNebulaSession(c.Nebula)
	if err != nil {
		glog.Fatal(err)
	}

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
		Nebula:     nebulaSession,
		PrivateKey: privateKey,
		JwtX:       middleware.NewJwtXMiddleware(redisClient, c).Handle,
		Captcha:    svcCaptcha,
		LLM: llm.Client{
			Config:     c.LLM,
			HTTPClient: httpClient,
		},
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

func (svcCtx *ServiceContext) GetUserFromCache(token string) (user m.User, err error) {

	cmd := svcCtx.Redis.Get(context.Background(), token)
	if cmd.Err() != nil {
		if cmd.Err() == redis.Nil {
			err = errors.New(http.StatusUnauthorized, "用户未登录")
		} else {
			err = cmd.Err()
		}

		glog.Error(err)
		return
	}

	userJSON, err := cmd.Bytes()
	if err != nil {
		glog.Error(err)
		return
	}

	err = json.Unmarshal(userJSON, &user)
	if err != nil {
		glog.Error(err)
		return
	}

	if user.ID == 0 {
		err = errors.New(http.StatusUnauthorized, "用戶未登录")
		glog.Error(err)
		return
	}

	return
}
