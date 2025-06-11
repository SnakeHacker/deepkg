package schema_ontology

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSchemaOntologyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSchemaOntologyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSchemaOntologyLogic {
	return &CreateSchemaOntologyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSchemaOntologyLogic) CreateSchemaOntology(req *types.CreateSchemaOntologyReq) error {
	// todo: add your logic here and delete this line

	return nil
}
