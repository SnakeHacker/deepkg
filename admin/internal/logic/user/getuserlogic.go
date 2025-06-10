package user

import (
	"context"
	"time"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.GetUserResp, err error) {
	userModel, err := dao.SelectUserByID(l.svcCtx.DB, req.ID)
	if err != nil {
		glog.Error(err)
		return
	}

	createdAt, _ := time.Parse(time.RFC3339, userModel.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, userModel.UpdatedAt)
	userModel.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
	userModel.UpdatedAt = updatedAt.Format("2006-01-02 15:04:05")

	resp = &types.GetUserResp{
		User: *userModel,
	}

	return
}
