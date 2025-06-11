package org

import (
	"context"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

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
	orgsModel, total, err := dao.SelectOrgs(l.svcCtx.DB, req.PageNumber, req.PageSize)
	if err != nil {
		glog.Error(err)
		return
	}

	orgs := []types.Organization{}
	for _, org := range orgsModel {
		orgs = append(orgs, types.Organization{
			ID:        int64(org.ID),
			OrgName:   org.OrgName,
			CreatedAt: org.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: org.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	resp = &types.GetOrgListResp{
		PageSize:      req.PageSize,
		PageNumber:    req.PageNumber,
		Organizations: orgs,
		Total:         total,
	}

	return
}
