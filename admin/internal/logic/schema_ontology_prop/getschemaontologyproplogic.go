package schema_ontology_prop

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/common"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSchemaOntologyPropLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSchemaOntologyPropLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSchemaOntologyPropLogic {
	return &GetSchemaOntologyPropLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSchemaOntologyPropLogic) GetSchemaOntologyProp(req *types.GetSchemaOntologyPropReq) (resp *types.GetSchemaOntologyPropResp, err error) {
	propModel, err := dao.SelectSchemaOntologyPropByID(l.svcCtx.DB, req.ID)
	if err != nil {
		glog.Error(err)
		return
	}

	userModel, err := dao.SelectUserByID(l.svcCtx.DB, int64(propModel.CreatorID))
	if err != nil {
		glog.Error(err)
		return
	}

	resp = &types.GetSchemaOntologyPropResp{
		SchemaOntologyProp: types.SchemaOntologyProp{
			ID:          int64(propModel.ID),
			PropName:    propModel.PropName,
			PropDesc:    propModel.PropDesc,
			OntologyID:  int64(propModel.OntologyID),
			WorkSpaceID: int64(propModel.WorkSpaceID),
			CreatorID:   int64(userModel.ID),
			CreatorName: userModel.Username,
			CreatedAt:   propModel.CreatedAt.Format(common.TIME_FORMAT),
		},
	}

	return
}
