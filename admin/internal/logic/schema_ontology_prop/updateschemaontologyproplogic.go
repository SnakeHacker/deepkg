package schema_ontology_prop

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSchemaOntologyPropLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSchemaOntologyPropLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSchemaOntologyPropLogic {
	return &UpdateSchemaOntologyPropLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSchemaOntologyPropLogic) UpdateSchemaOntologyProp(req *types.UpdateSchemaOntologyPropReq) (err error) {
	prop := req.SchemaOntologyProp
	propModel, err := dao.SelectSchemaOntologyPropByID(l.svcCtx.DB, prop.ID)
	if err != nil {
		glog.Error(err)
		return err
	}

	propModel.PropName = prop.PropName
	propModel.PropDesc = prop.PropDesc
	propModel.WorkSpaceID = int(prop.WorkSpaceID)
	propModel.CreatorID = int(prop.CreatorID)

	err = dao.UpdateSchemaOntologyProp(l.svcCtx.DB, &propModel)
	if err != nil {
		glog.Error(err)
		return err
	}

	return
}
