package schema_ontology

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/common"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSchemaOntologyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSchemaOntologyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSchemaOntologyLogic {
	return &GetSchemaOntologyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSchemaOntologyLogic) GetSchemaOntology(req *types.GetSchemaOntologyReq) (resp *types.GetSchemaOntologyResp, err error) {
	ontologyModel, err := dao.SelectSchemaOntologyByID(l.svcCtx.DB, req.ID)
	if err != nil {
		glog.Error(err)
		return
	}

	userModel, err := dao.SelectUserByID(l.svcCtx.DB, int64(ontologyModel.CreatorID))
	if err != nil {
		glog.Error(err)
		return
	}

	resp = &types.GetSchemaOntologyResp{
		SchemaOntology: types.SchemaOntology{
			ID:           int64(ontologyModel.ID),
			OntologyName: ontologyModel.OntologyName,
			OntologyDesc: ontologyModel.OntologyDesc,
			WorkSpaceID:  int64(ontologyModel.WorkSpaceID),
			CreatorID:    int64(userModel.ID),
			CreatorName:  userModel.Username,
			CreatedAt:    ontologyModel.CreatedAt.Format(common.TIME_FORMAT),
		},
	}

	return
}
