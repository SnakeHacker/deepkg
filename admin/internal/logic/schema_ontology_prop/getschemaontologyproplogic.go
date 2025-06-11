package schema_ontology_prop

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
