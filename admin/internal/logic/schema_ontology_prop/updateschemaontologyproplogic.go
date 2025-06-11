package schema_ontology_prop

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

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

func (l *UpdateSchemaOntologyPropLogic) UpdateSchemaOntologyProp(req *types.UpdateSchemaOntologyPropReq) error {
	// todo: add your logic here and delete this line

	return nil
}
