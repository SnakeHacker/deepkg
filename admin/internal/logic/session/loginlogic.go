package session

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/SnakeHacker/deepkg/admin/common"
	"github.com/SnakeHacker/deepkg/admin/common/captcha"
	"github.com/SnakeHacker/deepkg/admin/common/rsa"
	"github.com/SnakeHacker/deepkg/admin/common/werkzeug"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang-jwt/jwt"
	"github.com/golang/glog"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	"gorm.io/gorm"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if len(strings.TrimSpace(req.Account)) == 0 {
		return nil, errors.New(http.StatusBadRequest, "账号不能为空")
	}

	if len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New(http.StatusBadRequest, "密码不能为空")
	}

	if len(strings.TrimSpace(req.CaptchaID)) == 0 || len(strings.TrimSpace(req.CaptchaValue)) == 0 {
		return nil, errors.New(http.StatusBadRequest, "验证码参数错误")
	}

	successed, err := captcha.VerifyCaptcha(l.svcCtx.Captcha, req.CaptchaID, req.CaptchaValue)
	if err != nil || !successed {
		glog.Error(err)
		return
	}

	user, err := dao.SelectUserByAccount(l.svcCtx.DB, req.Account)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.New(http.StatusBadRequest, "用户名或密码不正确")
		}
		glog.Error(err)
		return
	}

	decryptedPassword, err := rsa.Decrypt(req.Password, l.svcCtx.PrivateKey)
	if err != nil {
		glog.Error(err)
		return
	}

	if correctPassword := werkzeug.CheckPasswordHash(decryptedPassword, user.PasswordHash); !correctPassword {
		err = errors.New(http.StatusBadRequest, "用户名或密码不正确")
		glog.Error(err)
		return
	}

	// 判断用户是否被禁用
	if user.Enable == common.USER_DISABLE {
		err = errors.New(http.StatusForbidden, "用户已被禁用")
		glog.Error(err)
		return
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, int64(user.ID))
	if err != nil {
		glog.Error(err)
		return
	}

	// 存入redis
	err = l.SetUserToCache(jwtToken, user, accessExpire)
	if err != nil {
		glog.Error(err)
		return
	}

	return &types.LoginResp{
		ID:           int64(user.ID),
		Username:     user.Username,
		Account:      user.Account,
		Role:         user.Role,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *LoginLogic) SetUserToCache(token string, user m.User, accessExpire int64) (err error) {
	// 将 User 对象转换为 JSON 格式
	userJSON, err := json.Marshal(user)
	if err != nil {
		glog.Error(err)
		return err
	}

	err = l.svcCtx.Redis.Set(context.Background(), token, userJSON, time.Duration(accessExpire)*time.Second).Err()
	if err != nil {
		glog.Error(err)
		return
	}
	return
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
