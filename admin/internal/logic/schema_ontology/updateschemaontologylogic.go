package schema_ontology

import (
	"context"

	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSchemaOntologyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSchemaOntologyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSchemaOntologyLogic {
	return &UpdateSchemaOntologyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSchemaOntologyLogic) UpdateSchemaOntology(req *types.UpdateSchemaOntologyReq) error {
	// todo: add your logic here and delete this line

	return nil
}
