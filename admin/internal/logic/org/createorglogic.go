package org

import (
	"context"
	"errors"
	"fmt"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/golang/glog"
	"gorm.io/gorm"

	m "github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

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
	result, err := dao.SelectOrgByName(l.svcCtx.DB, req.Organization.OrgName)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		glog.Error(err)
		return
	}

	if result.OrgName != "" {
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
