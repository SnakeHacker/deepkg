package schema_ontology

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
