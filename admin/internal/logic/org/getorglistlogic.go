package org

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrgListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrgListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrgListLogic {
	return &GetOrgListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrgListLogic) GetOrgList(req *types.GetOrgListReq) (resp *types.GetOrgListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
