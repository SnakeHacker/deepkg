package schema_ontology_prop

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSchemaOntologyPropLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSchemaOntologyPropLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSchemaOntologyPropLogic {
	return &CreateSchemaOntologyPropLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSchemaOntologyPropLogic) CreateSchemaOntologyProp(req *types.CreateSchemaOntologyPropReq) error {
	// todo: add your logic here and delete this line

	return nil
}
