package schema_ontology_prop

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSchemaOntologyPropsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSchemaOntologyPropsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSchemaOntologyPropsLogic {
	return &DeleteSchemaOntologyPropsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSchemaOntologyPropsLogic) DeleteSchemaOntologyProps(req *types.DeleteSchemaOntologyPropsReq) error {
	// todo: add your logic here and delete this line

	return nil
}
