package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SnakeHacker/deepkg/admin/internal/config"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/x/errors"

	"github.com/golang/glog"
)

type JwtXMiddleware struct {
	Redis  redis.UniversalClient
	Config config.Config
}

func NewJwtXMiddleware(redis redis.UniversalClient, c config.Config) *JwtXMiddleware {
	return &JwtXMiddleware{
		Redis:  redis,
		Config: c,
	}
}

func (m *JwtXMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != m.Config.Auth.HardAccessSecret {

			user := gorm_model.User{}
			cmd := m.Redis.Get(context.Background(), token)
			var err error
			if cmd.Err() != nil {
				if cmd.Err() == redis.Nil {
					err = errors.New(http.StatusUnauthorized, "用户未登录")
				} else {
					err = cmd.Err()
					glog.Error(err)
				}
				// 返回未授权响应
				glog.Errorf("%v,%v", err, r.URL)
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, err)
				return
			}

			userJSON, err := cmd.Bytes()
			if err != nil {
				// 获取用户信息失败，记录日志并返回错误
				glog.Error(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "Failed to get user data from Redis (JwtMiddleware)")
				return
			}

			// 将获取到的用户信息扫描到 user 对象中
			err = json.Unmarshal(userJSON, &user)
			if err != nil {
				// 解析用户信息失败，记录日志并返回错误
				glog.Error(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "Failed to parse user data (JwtMiddleware)")
				return
			}

			if user.ID == 0 {
				err = errors.New(http.StatusUnauthorized, "用戶未登录")
				glog.Errorf("%v,%v", err, r.URL)

				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, err)
				return
			}
		}

		next(w, r)
	}
}
