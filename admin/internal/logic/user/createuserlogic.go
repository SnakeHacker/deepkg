package user

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/common"
	"github.com/SnakeHacker/deepkg/admin/common/werkzeug"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (err error) {
	// 检查 org 是否存在
	_, err = dao.SelectOrgByID(l.svcCtx.DB, int(req.User.OrgID))
	if err != nil {
		glog.Error(err)
		return
	}

	// rsa解密
	decryptedPassword, err := common.PasswordValidate(req.User.Password, l.svcCtx.PrivateKey)
	if err != nil {
		glog.Error(err)
		return
	}

	user := &m.User{
		UserCode:     uuid.New().String(),
		OrgID:        int(req.User.OrgID),
		Account:      req.User.Account,
		Username:     req.User.Username,
		PasswordHash: werkzeug.GeneratePasswordHash(decryptedPassword),
		Phone:        req.User.Phone,
		Mail:         req.User.Mail,
		Enable:       req.User.Enable,
		Role:         req.User.Role,
		Avatar:       req.User.Avatar,
	}

	err = dao.CreateUser(l.svcCtx.DB, user)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
