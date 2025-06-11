package schema_ontology_prop

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSchemaOntologyPropListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSchemaOntologyPropListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSchemaOntologyPropListLogic {
	return &GetSchemaOntologyPropListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSchemaOntologyPropListLogic) GetSchemaOntologyPropList(req *types.GetSchemaOntologyPropListReq) (resp *types.GetSchemaOntologyPropListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
