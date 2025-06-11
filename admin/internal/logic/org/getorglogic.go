package org

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/common"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

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
	orgModel, err := dao.SelectOrgByID(l.svcCtx.DB, req.ID)
	if err != nil {
		glog.Error(err)
		return
	}

	result := types.Organization{
		ID:        int64(orgModel.ID),
		OrgName:   orgModel.OrgName,
		CreatedAt: orgModel.CreatedAt.Format(common.TIME_FORMAT),
		UpdatedAt: orgModel.UpdatedAt.Format(common.TIME_FORMAT),
	}

	resp = &types.GetOrgResp{
		Organization: result,
	}

	return
}
