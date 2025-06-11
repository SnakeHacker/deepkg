package schema_ontology

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSchemaOntologyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSchemaOntologyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSchemaOntologyListLogic {
	return &GetSchemaOntologyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSchemaOntologyListLogic) GetSchemaOntologyList(req *types.GetSchemaOntologyListReq) (resp *types.GetSchemaOntologyListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
