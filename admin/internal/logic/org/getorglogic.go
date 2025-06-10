package org

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrgLogic {
	return &GetOrgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrgLogic) GetOrg(req *types.GetOrgReq) (resp *types.GetOrgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
