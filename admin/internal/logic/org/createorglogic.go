package org

import (
	"context"
	"fmt"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrgLogic {
	return &CreateOrgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrgLogic) CreateOrg(req *types.CreateOrgReq) (err error) {
	// 检查组织名称是否已存在
	exists, err := dao.CheckOrgNameExists(l.svcCtx.DB, req.Organization.OrgName)

	if err != nil {
		glog.Error(err)
		return
	}

	if exists {
		err = fmt.Errorf("organization with name %s already exists", req.Organization.OrgName)
		glog.Error(err)
		return
	}

	org := &m.Organization{
		OrgName: req.Organization.OrgName,
	}

	err = dao.CreateOrg(l.svcCtx.DB, org)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
