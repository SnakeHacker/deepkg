package user

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (err error) {
	userModel, err := dao.SelectUserModleByID(l.svcCtx.DB, req.User.ID)
	if err != nil {
		glog.Error(err)
		return
	}
	_, err = dao.SelectOrgByID(l.svcCtx.DB, int(req.User.OrgID))
	if err != nil {
		glog.Error(err)
		return
	}

	// 更新字段
	if req.User.OrgID != 0 {
		userModel.OrgID = int(req.User.OrgID)
	}
	if req.User.Username != "" {
		userModel.Username = req.User.Username
	}
	if req.User.Phone != "" {
		userModel.Phone = req.User.Phone
	}
	if req.User.Mail != "" {
		userModel.Mail = req.User.Mail
	}
	if req.User.Enable != 0 {
		userModel.Enable = req.User.Enable
	}
	if req.User.Avatar != "" {
		userModel.Avatar = req.User.Avatar
	}

	// 保存更新
	err = dao.UpdateUser(l.svcCtx.DB, userModel)
	if err != nil {
		glog.Error(err)
		return err
	}

	return
}
