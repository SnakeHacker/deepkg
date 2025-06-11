package org

import (
	"context"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/golang/glog"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrgLogic {
	return &UpdateOrgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOrgLogic) UpdateOrg(req *types.UpdateOrgReq) (err error) {
	orgModel, err := dao.SelectOrgByID(l.svcCtx.DB, req.Organization.ID)
	if err != nil {
		glog.Error(err)
		return
	}

	// 更新字段
	if req.Organization.OrgName != "" {
		orgModel.OrgName = req.Organization.OrgName
	}

	err = dao.UpdateOrg(l.svcCtx.DB, orgModel)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
