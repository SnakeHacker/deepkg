package org

import (
	"context"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/golang/glog"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOrgsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOrgsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrgsLogic {
	return &DeleteOrgsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteOrgsLogic) DeleteOrgs(req *types.DeleteOrgsReq) (err error) {
	err = dao.DeleteOrgsByIDs(l.svcCtx.DB, req.Ids)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
